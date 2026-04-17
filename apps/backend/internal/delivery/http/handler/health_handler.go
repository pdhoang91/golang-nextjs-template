package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/your-org/fullstack-template/apps/backend/internal/config"
)

type HealthHandler struct {
	config *config.Config
}

func NewHealthHandler(config *config.Config) *HealthHandler {
	return &HealthHandler{
		config: config,
	}
}

func (h *HealthHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":      "ok",
		"service":     h.config.AppName,
		"environment": h.config.AppEnv,
		"timestamp":   time.Now().UTC(),
	})
}
