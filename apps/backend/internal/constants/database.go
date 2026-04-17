package constants

const (
	PostgresDSNFormat          = "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	MigrationDatabaseURLFormat = "postgres://%s:%s@%s:%d/%s?sslmode=%s"
	FileSchemeFormat           = "file://%s"
)
