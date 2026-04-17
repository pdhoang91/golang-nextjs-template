MIGRATE_ACTION_UP ?= up
MIGRATE_ACTION_DOWN ?= down

.PHONY: run test migrate-up migrate-down fmt

run:
	go run ./cmd/server

test:
	go test ./...

migrate-up:
	go run ./cmd/migrate $(MIGRATE_ACTION_UP)

migrate-down:
	go run ./cmd/migrate $(MIGRATE_ACTION_DOWN)

fmt:
	gofmt -w ./cmd ./internal
