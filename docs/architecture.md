# Architecture Notes

## Backend

Backend đi theo Clean Architecture tối giản:

- **domain**: entity và business error
- **usecase**: orchestration + validation + business rule
- **repository**: interface data access
- **infrastructure/persistence**: implementation cụ thể bằng GORM/Postgres
- **delivery/http**: Gin handlers, DTO, routes, middleware
- **bootstrap**: wiring dependencies
- **config**: load env config

### Dependency direction

```text
handler -> usecase -> repository interface <- repository implementation
```

Điều quan trọng là:

- `domain` không biết Gin/GORM
- `usecase` không biết HTTP request/response
- `repository implementation` không chứa business logic

## Frontend

Frontend theo App Router + feature-based structure:

- `app/`: route entry
- `features/`: logic và component theo từng feature
- `services/`: API access layer dùng chung
- `types/`: shared typing
- `hooks/`: reusable hook dùng chung
- `components/`: UI/shared layout

### Data flow

```text
page -> feature component -> feature hook -> service -> api client -> backend
```

Cách này giúp:

- page mỏng
- logic fetch không nhồi vào page
- UI component dễ tách / test / thay đổi
