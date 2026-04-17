package router

import (
	"log/slog"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/your-org/fullstack-template/apps/backend/internal/config"
	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/middleware"
)

type RouteRegistrar interface {
	RegisterRoutes(rg *gin.RouterGroup)
}

func New(cfg *config.Config, logger *slog.Logger, registrars []RouteRegistrar) *gin.Engine {
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
		AllowHeaders:     []string{constants.HeaderOrigin, constants.HeaderContentType, constants.HeaderAccept, constants.HeaderAuthorization, constants.HeaderRequestID},
		ExposeHeaders:    []string{constants.HeaderRequestID},
		AllowCredentials: true,
	}))

	apiV1 := r.Group(constants.APIV1Prefix)

	for _, registrar := range registrars {
		if registrar != nil {
			registrar.RegisterRoutes(apiV1)
		}
	}

	return r
}
