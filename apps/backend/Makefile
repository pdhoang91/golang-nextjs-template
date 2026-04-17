.PHONY: run test migrate-up migrate-down fmt

run:
	go run ./cmd/server

test:
	go test ./...

migrate-up:
	go run ./cmd/migrate up

migrate-down:
	go run ./cmd/migrate down

fmt:
	gofmt -w ./cmd ./internal
