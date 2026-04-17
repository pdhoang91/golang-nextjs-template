# API App

Backend API dùng Go + Gin + GORM + PostgreSQL.

## Kiến trúc

```text
cmd/
internal/
  bootstrap/
  config/
  domain/
  usecase/<module>/
  infrastructure/
  delivery/http/
migrations/
```

### Tư duy chính

- **domain**: entity + business errors + repository contract
- **usecase/<module>**: business rule
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

- `GET /api/v1/health`
- `GET /api/v1/todos`
- `POST /api/v1/todos`

## Chạy migration

```bash
make migrate-up
make migrate-down
```

## Test

```bash
go test ./...
```

## Cách thêm module mới

Ví dụ module `users`.

## Ví dụ module không cần repository

`health` là ví dụ module chỉ cần:
- `internal/domain/health/health.go`
- `internal/usecase/health/usecase.go`
- `internal/delivery/http/dto/health_response.go`
- `internal/delivery/http/handler/health_handler.go`

Module kiểu này phù hợp cho các query đơn giản, check status, hoặc logic chỉ ghép dữ liệu từ config/runtime.

### 1. Tạo domain
- `internal/domain/user/user.go`

### 2. Tạo repository interface
- `internal/domain/user/repository.go`

### 3. Tạo usecase
- `internal/usecase/user/usecase.go`

### 4. Tạo repository implementation
- `internal/infrastructure/persistence/postgres/user_repository.go`

### 5. Tạo DTO + handler
- `internal/delivery/http/dto/user.go`
- `internal/delivery/http/handler/user_handler.go`

### 6. Đăng ký ở bootstrap
Khởi tạo repo, usecase, handler trong `internal/bootstrap/app.go`.

### 7. Đăng ký route
Cho handler tự đăng ký route và truyền registrar vào `internal/delivery/http/router/router.go`.

### 8. Tạo migration
Thêm file trong `migrations/`.

## Những điều nên tránh

- không parse business rule trong handler
- không import Gin/GORM vào domain
- không viết repository interface trong infrastructure
- không tạo shared package quá sớm
