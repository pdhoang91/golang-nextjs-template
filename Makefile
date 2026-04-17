PROJECT_NAME=fullstack-template

.PHONY: help up down logs build backend-dev frontend-dev backend-test migrate-up migrate-down format

help:
	@echo "make up            - start all containers"
	@echo "make down          - stop all containers"
	@echo "make logs          - follow docker compose logs"
	@echo "make build         - rebuild docker images"
	@echo "make backend-dev   - run backend locally"
	@echo "make frontend-dev  - run frontend locally"
	@echo "make backend-test  - run backend tests"
	@echo "make migrate-up    - run backend migrations locally"
	@echo "make migrate-down  - rollback one backend migration locally"
	@echo "make format        - format backend code"

up:
	docker compose up --build

down:
	docker compose down -v

logs:
	docker compose logs -f

build:
	docker compose build

backend-dev:
	cd apps/backend && go run ./cmd/server

frontend-dev:
	cd apps/frontend && npm install && npm run dev

backend-test:
	cd apps/backend && go test ./...

migrate-up:
	cd apps/backend && go run ./cmd/migrate up

migrate-down:
	cd apps/backend && go run ./cmd/migrate down

format:
	cd apps/backend && gofmt -w ./cmd ./internal
