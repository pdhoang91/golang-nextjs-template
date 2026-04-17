package bootstrap

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/your-org/fullstack-template/apps/backend/internal/config"
	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/handler"
	"github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/router"
	"github.com/your-org/fullstack-template/apps/backend/internal/infrastructure/database"
	infraLogger "github.com/your-org/fullstack-template/apps/backend/internal/infrastructure/logger"
	"github.com/your-org/fullstack-template/apps/backend/internal/infrastructure/migration"
	postgresrepo "github.com/your-org/fullstack-template/apps/backend/internal/infrastructure/persistence/postgres"
	healthusecase "github.com/your-org/fullstack-template/apps/backend/internal/usecase/health"
	usetodo "github.com/your-org/fullstack-template/apps/backend/internal/usecase/todo"
	"gorm.io/gorm"
)

type App struct {
	cfg        *config.Config
	logger     *slog.Logger
	httpServer *http.Server
	database   *gorm.DB
}

func NewApp(_ context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	logger := infraLogger.New()

	if err := migration.Up(cfg.MigrationDatabaseURL(), cfg.MigrationsPath); err != nil {
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	db, err := database.NewPostgres(cfg)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	healthUsecase := healthusecase.NewHealthUsecase(cfg.AppName, cfg.AppEnv)
	todoRepository := postgresrepo.NewTodoRepository(db)
	todoUsecase := usetodo.NewTodoUsecase(todoRepository)

	healthHandler := handler.NewHealthHandler(healthUsecase)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r := router.New(cfg, logger, []router.RouteRegistrar{
		healthHandler,
		todoHandler,
	})

	httpServer := &http.Server{
		Addr:              cfg.Address(),
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &App{
		cfg:        cfg,
		logger:     logger,
		httpServer: httpServer,
		database:   db,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	errorCh := make(chan error, 1)

	go func() {
		a.logger.Info(constants.LogStartingHTTPServer, "address", a.cfg.Address())
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errorCh <- err
			return
		}
		errorCh <- nil
	}()

	select {
	case err := <-errorCh:
		return err
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(a.cfg.ShutdownTimeout)*time.Second)
		defer cancel()

		a.logger.Info(constants.LogShuttingDownApplication)

		if err := a.httpServer.Shutdown(shutdownCtx); err != nil {
			return err
		}

		sqlDB, err := a.database.DB()
		if err == nil {
			_ = sqlDB.Close()
		}

		return nil
	}
}

func (a *App) Logger() *slog.Logger {
	return a.logger
}
