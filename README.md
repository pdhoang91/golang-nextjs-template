# Fullstack Monorepo Template

Production-ready starter template cho **Go + Next.js + PostgreSQL + Docker Compose** theo hướng **clean, dễ maintain, dễ mở rộng**.

## 1. Kiến trúc tổng thể

Template này chia theo monorepo:

- `apps/backend`: backend Golang theo Clean Architecture
- `apps/frontend`: frontend Next.js App Router theo feature-based structure
- `deployments/docker`: Dockerfile riêng cho từng app
- `docs`: tài liệu kiến trúc và hướng dẫn mở rộng
- root: `docker-compose.yml`, `.env.example`, `Makefile`

### Backend stack

- Go 1.24
- Gin
- GORM
- PostgreSQL
- golang-migrate
- slog (structured logging)

### Frontend stack

- Next.js App Router
- TypeScript
- ESLint + Prettier
- Shared API client
- Feature-based structure hợp lý

## 2. Vì sao chọn kiến trúc này

### Gin
Gin đủ nhẹ, phổ biến, middleware ecosystem tốt, dễ onboarding cho team, phù hợp template CRUD/API lâu dài.

### GORM
GORM giúp template khởi động nhanh, ít boilerplate, phù hợp base project dùng lại nhiều lần. Clean Architecture vẫn được giữ vì GORM chỉ xuất hiện ở infrastructure/repository implementation.  
Khi hệ thống nhiều query phức tạp hơn, có thể chuyển một số module sang `sqlc`.

### Clean Architecture ở backend
Luồng phụ thuộc:

`delivery -> usecase -> repository interface -> repository implementation`

- handler không giữ business logic
- usecase giữ business rule
- repository chỉ làm data access
- implementation không rò rỉ lên domain/usecase

### Feature-based ở frontend
Tách theo feature giúp scale tốt hơn kiểu gom toàn bộ page/component/service vào một chỗ.  
Tuy nhiên vẫn giữ `components/`, `hooks/`, `services/`, `types/` dùng chung để không bị quá nặng.

## 3. Cây thư mục

```text
fullstack-template/
├── .env.example
├── .gitignore
├── Makefile
├── README.md
├── docker-compose.yml
├── apps/
│   ├── backend/
│   │   ├── .env.example
│   │   ├── Makefile
│   │   ├── README.md
│   │   ├── go.mod
│   │   ├── cmd/
│   │   │   ├── migrate/
│   │   │   │   └── main.go
│   │   │   └── server/
│   │   │       └── main.go
│   │   ├── internal/
│   │   │   ├── bootstrap/
│   │   │   │   └── app.go
│   │   │   ├── config/
│   │   │   │   └── config.go
│   │   │   ├── delivery/
│   │   │   │   └── http/
│   │   │   │       ├── dto/
│   │   │   │       │   └── todo.go
│   │   │   │       ├── handler/
│   │   │   │       │   ├── health_handler.go
│   │   │   │       │   └── todo_handler.go
│   │   │   │       ├── middleware/
│   │   │   │       │   ├── auth.go
│   │   │   │       │   ├── logger.go
│   │   │   │       │   └── request_id.go
│   │   │   │       ├── response/
│   │   │   │       │   └── response.go
│   │   │   │       └── router/
│   │   │   │           └── router.go
│   │   │   ├── domain/
│   │   │   │   └── todo/
│   │   │   │       └── todo.go
│   │   │   ├── infrastructure/
│   │   │   │   ├── database/
│   │   │   │   │   └── postgres.go
│   │   │   │   ├── logger/
│   │   │   │   │   └── logger.go
│   │   │   │   ├── migration/
│   │   │   │   │   └── runner.go
│   │   │   │   └── persistence/
│   │   │   │       └── postgres/
│   │   │   │           └── todo_repository.go
│   │   │   ├── repository/
│   │   │   │   └── todo_repository.go
│   │   │   └── usecase/
│   │   │       ├── todo_usecase.go
│   │   │       └── todo_usecase_test.go
│   │   └── migrations/
│   │       ├── 000001_create_todos.down.sql
│   │       └── 000001_create_todos.up.sql
│   └── frontend/
│       ├── .env.example
│       ├── .eslintrc.json
│       ├── .prettierrc
│       ├── README.md
│       ├── next-env.d.ts
│       ├── next.config.ts
│       ├── package.json
│       ├── tsconfig.json
│       ├── app/
│       │   ├── globals.css
│       │   ├── layout.tsx
│       │   └── page.tsx
│       ├── components/
│       │   ├── layout/
│       │   │   └── container.tsx
│       │   └── ui/
│       │       ├── card.tsx
│       │       └── status-badge.tsx
│       ├── config/
│       │   └── env.ts
│       ├── constants/
│       │   └── routes.ts
│       ├── features/
│       │   ├── health/
│       │   │   ├── components/
│       │   │   │   └── health-status-card.tsx
│       │   │   └── hooks/
│       │   │       └── use-health.ts
│       │   └── todos/
│       │       ├── components/
│       │       │   ├── todo-create-form.tsx
│       │       │   └── todo-list-card.tsx
│       │       └── hooks/
│       │           └── use-todos.ts
│       ├── hooks/
│       │   └── use-async-state.ts
│       ├── lib/
│       │   └── format-date.ts
│       ├── services/
│       │   ├── api-client.ts
│       │   ├── health.service.ts
│       │   └── todo.service.ts
│       └── types/
│           ├── api.ts
│           ├── health.ts
│           └── todo.ts
├── deployments/
│   └── docker/
│       ├── backend/
│       │   └── Dockerfile
│       └── frontend/
│           └── Dockerfile
└── docs/
    ├── adding-module.md
    └── architecture.md
```

## 4. Chạy bằng Docker Compose

```bash
cp .env.example .env
docker compose up --build
```

Sau đó truy cập:

- Frontend: `http://localhost:3000`
- Backend health: `http://localhost:8080/health`
- Backend todos: `http://localhost:8080/api/v1/todos`

## 5. Chạy local từng app

### Backend
```bash
cp apps/backend/.env.example apps/backend/.env
cd apps/backend
go mod tidy
go run ./cmd/server
```

### Frontend
```bash
cp apps/frontend/.env.example apps/frontend/.env.local
cd apps/frontend
npm install
npm run dev
```

## 6. Hướng dẫn nhanh thêm module mới

Xem file:

- `docs/adding-module.md`
- `apps/backend/README.md`

Pattern backend cho module mới:

1. `internal/domain/<module>`
2. `internal/repository/<module>_repository.go`
3. `internal/usecase/<module>_usecase.go`
4. `internal/infrastructure/persistence/postgres/<module>_repository.go`
5. `internal/delivery/http/handler/<module>_handler.go`
6. `internal/delivery/http/dto/<module>.go`
7. đăng ký route trong `router/router.go`
8. thêm migration SQL

## 7. Các điểm cần tránh

- nhét business logic vào handler
- để repository quyết định rule nghiệp vụ
- để domain phụ thuộc framework hoặc ORM
- viết shared package quá sớm khi chưa thực sự tái sử dụng
- tạo abstraction thừa khi mới có 1 implementation

## 8. Các cải tiến nên làm ở phiên bản tiếp theo

- authentication / authorization
- Redis cache
- background job / queue
- observability: metrics, tracing, log shipping
- CI/CD pipeline
- Swagger / OpenAPI
- RBAC
- file storage (S3/MinIO)
- refresh token / session management
- pagination / filtering / sorting helper
- graceful config profiles cho dev/staging/prod
