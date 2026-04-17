package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
	httpdto "github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/dto"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/middleware"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/response"
	usetodo "github.com/your-org/fullstack-template/apps/backend/internal/usecase/todo"
)

type TodoHandler struct {
	todoUsecase usetodo.TodoUsecase
}

func NewTodoHandler(todoUsecase usetodo.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (h *TodoHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET(constants.TodosRoute, h.List)
	rg.POST(constants.TodosRoute, middleware.AuthPlaceholder(), h.Create)
}

func (h *TodoHandler) List(c *gin.Context) {
	items, err := h.todoUsecase.List(c.Request.Context())
	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, constants.ErrFailedToFetchTodos)
		return
	}

	response.WriteSuccess(c, http.StatusOK, httpdto.ToTodoResponses(items))
}

func (h *TodoHandler) Create(c *gin.Context) {
	var req httpdto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.WriteError(c, http.StatusBadRequest, constants.ErrInvalidRequestPayload)
		return
	}

	created, err := h.todoUsecase.Create(c.Request.Context(), usetodo.CreateTodoInput{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		response.WriteMappedError(c, err, http.StatusInternalServerError, constants.ErrFailedToCreateTodo)
		return
	}

	response.WriteSuccess(c, http.StatusCreated, httpdto.ToTodoResponse(created))
}
