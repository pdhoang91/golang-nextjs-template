package router

import (
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/your-org/fullstack-template/apps/backend/internal/config"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/handler"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/middleware"
)

type Handlers struct {
	Health *handler.HealthHandler
	Todo   *handler.TodoHandler
}

func New(cfg *config.Config, logger *slog.Logger, handlers Handlers) *gin.Engine {
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	r.Use(middleware.RequestLogger(logger))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins(),
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Request-ID"},
		ExposeHeaders:    []string{"X-Request-ID"},
		AllowCredentials: true,
	}))

	r.GET("/health", handlers.Health.Get)

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/todos", handlers.Todo.List)
		apiV1.POST("/todos", middleware.AuthPlaceholder(), handlers.Todo.Create)
	}

	return r
}
