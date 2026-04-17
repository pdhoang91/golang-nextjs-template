package main

import (
	"fmt"
	"log"
	"os"

	"github.com/your-org/fullstack-template/apps/backend/internal/config"
	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
	"github.com/your-org/fullstack-template/apps/backend/internal/infrastructure/migration"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf(constants.MigrationUsageFormat, constants.MigrationActionUp, constants.MigrationActionDown)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	direction := os.Args[1]

	switch direction {
	case constants.MigrationActionUp:
		err = migration.Up(cfg.MigrationDatabaseURL(), cfg.MigrationsPath)
	case constants.MigrationActionDown:
		err = migration.Down(cfg.MigrationDatabaseURL(), cfg.MigrationsPath)
	default:
		log.Fatalf(constants.MigrationUnsupportedCommandFormat, direction)
	}

	if err != nil {
		log.Fatalf(constants.MigrationFailedFormat, direction, err)
	}

	fmt.Printf(constants.MigrationCompletedFormat+"\n", direction)
}
