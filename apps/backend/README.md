# API App

Backend API dùng Go + Gin + GORM + PostgreSQL.

## Kiến trúc

```text
cmd/
internal/
  bootstrap/
  config/
  domain/
  repository/
  usecase/
  infrastructure/
  delivery/http/
migrations/
```

### Tư duy chính

- **domain**: entity + business errors
- **usecase**: business rule
- **repository**: contract
- **repository implementation**: GORM/Postgres
- **delivery/http**: handler, DTO, route, middleware
- **bootstrap**: dependency wiring

## Chạy local

```bash
cp .env.example .env
go mod tidy
go run ./cmd/server
```

API mặc định:

- `GET /health`
- `GET /api/v1/todos`
- `POST /api/v1/todos`

## Chạy migration

```bash
go run ./cmd/migrate up
go run ./cmd/migrate down
```

## Test

```bash
go test ./...
```

## Cách thêm module mới

Ví dụ module `users`.

### 1. Tạo domain
- `internal/domain/user/user.go`

### 2. Tạo repository interface
- `internal/repository/user_repository.go`

### 3. Tạo usecase
- `internal/usecase/user_usecase.go`

### 4. Tạo repository implementation
- `internal/infrastructure/persistence/postgres/user_repository.go`

### 5. Tạo DTO + handler
- `internal/delivery/http/dto/user.go`
- `internal/delivery/http/handler/user_handler.go`

### 6. Đăng ký ở bootstrap
Khởi tạo repo, usecase, handler trong `internal/bootstrap/app.go`.

### 7. Đăng ký route
Thêm route ở `internal/delivery/http/router/router.go`.

### 8. Tạo migration
Thêm file trong `migrations/`.

## Những điều nên tránh

- không parse business rule trong handler
- không import Gin/GORM vào domain
- không viết repository interface trong infrastructure
- không tạo shared package quá sớm
