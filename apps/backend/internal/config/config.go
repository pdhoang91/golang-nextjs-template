package config

import (
	"fmt"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
)

type Config struct {
	AppName          string `env:"APP_NAME" envDefault:"template-backend"`
	AppEnv           string `env:"APP_ENV" envDefault:"development"`
	AppPort          string `env:"APP_PORT" envDefault:"8080"`
	ShutdownTimeout  int    `env:"APP_SHUTDOWN_TIMEOUT" envDefault:"10"`
	MigrationsPath   string `env:"MIGRATIONS_PATH" envDefault:"migrations"`
	DBHost           string `env:"DB_HOST" envDefault:"localhost"`
	DBPort           int    `env:"DB_PORT" envDefault:"5432"`
	DBName           string `env:"DB_NAME" envDefault:"app_db"`
	DBUser           string `env:"DB_USER" envDefault:"postgres"`
	DBPassword       string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBSSLMode        string `env:"DB_SSLMODE" envDefault:"disable"`
	DBMaxOpenConns   int    `env:"DB_MAX_OPEN_CONNS" envDefault:"25"`
	DBMaxIdleConns   int    `env:"DB_MAX_IDLE_CONNS" envDefault:"10"`
	DBConnMaxLifeMin int    `env:"DB_CONN_MAX_LIFETIME_MINUTES" envDefault:"30"`
	CORSOrigins      string `env:"CORS_ALLOWED_ORIGINS" envDefault:"http://localhost:3000"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c Config) Address() string {
	return ":" + c.AppPort
}

func (c Config) AllowedOrigins() []string {
	parts := strings.Split(c.CORSOrigins, ",")
	origins := make([]string, 0, len(parts))

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			origins = append(origins, trimmed)
		}
	}

	return origins
}

func (c Config) DatabaseDSN() string {
	return fmt.Sprintf(
		constants.PostgresDSNFormat,
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBPassword,
		c.DBName,
		c.DBSSLMode,
	)
}

func (c Config) MigrationDatabaseURL() string {
	return fmt.Sprintf(
		constants.MigrationDatabaseURLFormat,
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}
