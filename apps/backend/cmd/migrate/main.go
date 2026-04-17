package main

import (
	"fmt"
	"log"
	"os"

	"github.com/your-org/fullstack-template/apps/backend/internal/config"
	"github.com/your-org/fullstack-template/apps/backend/internal/infrastructure/migration"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./cmd/migrate [up|down]")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	direction := os.Args[1]

	switch direction {
	case "up":
		err = migration.Up(cfg.MigrationDatabaseURL(), cfg.MigrationsPath)
	case "down":
		err = migration.Down(cfg.MigrationDatabaseURL(), cfg.MigrationsPath)
	default:
		log.Fatalf("unsupported migration command: %s", direction)
	}

	if err != nil {
		log.Fatalf("migration %s failed: %v", direction, err)
	}

	fmt.Printf("migration %s completed\n", direction)
}
