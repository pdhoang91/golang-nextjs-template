# Adding a new backend module

Ví dụ muốn thêm module `users`.

## 1. Domain

Tạo:

- `internal/domain/user/user.go`

Chỉ chứa entity, value object, error nghiệp vụ cơ bản.

## 2. Repository interface

Tạo:

- `internal/domain/user/repository.go`

Định nghĩa các hành vi cần cho usecase.

## 3. Usecase

Tạo:

- `internal/usecase/user/usecase.go`

Usecase nhận repository interface qua constructor.

## 4. Repository implementation

Tạo:

- `internal/infrastructure/persistence/postgres/user_repository.go`

Đây là nơi dùng GORM/Postgres model mapping.

Nếu module không cần truy cập database, có thể bỏ qua bước này. Ví dụ `health` chỉ cần domain + usecase + delivery.

## 5. HTTP delivery

Tạo:

- `internal/delivery/http/dto/user.go`
- `internal/delivery/http/handler/user_handler.go`

Handler chỉ:
- parse request
- gọi usecase
- map response

## 6. Register dependency

Cập nhật `internal/bootstrap/app.go`:
- khởi tạo repo
- khởi tạo usecase
- khởi tạo handler

## 7. Register routes

Các handler tự đăng ký route, rồi đăng ký chúng trong `internal/delivery/http/router/router.go`

## 8. Create migration

Tạo migration mới trong `migrations/`.

## 9. Add tests

Ít nhất nên có:
- usecase test
- handler contract test hoặc integration test nhỏ
