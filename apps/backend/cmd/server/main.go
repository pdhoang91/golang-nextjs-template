package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/your-org/fullstack-template/apps/backend/internal/bootstrap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app, err := bootstrap.NewApp(ctx)
	if err != nil {
		slog.Error("failed to bootstrap application", "error", err)
		os.Exit(1)
	}

	if err := app.Run(ctx); err != nil {
		app.Logger().Error("application stopped with error", "error", err)
		os.Exit(1)
	}
}
