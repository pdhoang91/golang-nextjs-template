package constants

const (
	MigrationActionUp   = "up"
	MigrationActionDown = "down"

	MigrationUsageFormat              = "usage: go run ./cmd/migrate [%s|%s]"
	MigrationUnsupportedCommandFormat = "unsupported migration command: %s"
	MigrationFailedFormat             = "migration %s failed: %v"
	MigrationCompletedFormat          = "migration %s completed"
)
