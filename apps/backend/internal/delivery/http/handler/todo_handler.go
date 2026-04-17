package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	httpdto "github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/dto"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/response"
	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
	"github.com/your-org/fullstack-template/apps/backend/internal/usecase"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (h *TodoHandler) List(c *gin.Context) {
	items, err := h.todoUsecase.List(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to fetch todos")
		return
	}

	response.Success(c, http.StatusOK, httpdto.ToTodoResponses(items))
}

func (h *TodoHandler) Create(c *gin.Context) {
	var req httpdto.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request payload")
		return
	}

	created, err := h.todoUsecase.Create(c.Request.Context(), usecase.CreateTodoInput{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		if errors.Is(err, todo.ErrInvalidTitle) {
			response.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		response.Error(c, http.StatusInternalServerError, "failed to create todo")
		return
	}

	response.Success(c, http.StatusCreated, httpdto.ToTodoResponse(created))
}
