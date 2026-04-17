package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
	httpdto "github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/dto"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/response"
	healthusecase "github.com/your-org/fullstack-template/apps/backend/internal/usecase/health"
)

type HealthHandler struct {
	healthUsecase healthusecase.HealthUsecase
}

func NewHealthHandler(healthUsecase healthusecase.HealthUsecase) *HealthHandler {
	return &HealthHandler{
		healthUsecase: healthUsecase,
	}
}

func (h *HealthHandler) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET(constants.HealthRoute, h.Get)
}

func (h *HealthHandler) Get(c *gin.Context) {
	status, err := h.healthUsecase.Get(c.Request.Context())
	if err != nil {
		response.WriteError(c, http.StatusInternalServerError, constants.ErrFailedToFetchHealth)
		return
	}

	response.WriteSuccess(c, http.StatusOK, httpdto.ToHealthResponse(status))
}
