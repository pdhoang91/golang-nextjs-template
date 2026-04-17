package migration

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
)

func Up(databaseURL string, migrationsPath string) error {
	m, err := migrate.New(fmt.Sprintf(constants.FileSchemeFormat, migrationsPath), databaseURL)
	if err != nil {
		return err
	}
	defer func() { _, _ = m.Close() }()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func Down(databaseURL string, migrationsPath string) error {
	m, err := migrate.New(fmt.Sprintf(constants.FileSchemeFormat, migrationsPath), databaseURL)
	if err != nil {
		return err
	}
	defer func() { _, _ = m.Close() }()

	if err := m.Steps(-1); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
