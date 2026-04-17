# Project Rules

## Scope

- This repo is a monorepo with:
- `apps/backend`: Go API
- `apps/frontend`: Next.js app
- `deployments/docker`: container build files
- `docs`: architecture and module guides

- Source of truth for repo structure and extension patterns:
- `README.md`
- `apps/backend/README.md`
- `apps/frontend/README.md`
- `docs/adding-module.md`

## Repo-Wide Rules

- Keep changes aligned with the current structure. Do not introduce paths or conventions from older layouts.
- Reuse existing abstractions before adding new ones.
- Prefer incremental refactors over broad rewrites.
- Add or update tests when behavior changes.
- Keep functions, types, and packages focused on one responsibility.
- Favor explicit wiring over hidden magic.
- Prefer constants for repeated literals and technical strings such as routes, headers, log keys, statuses, error codes, DSN formats, and migration actions.
- Place constants at the layer or feature that owns the value.
- Keep one-off copy, local test data, and truly isolated literals inline when extracting them would not improve reuse or readability.

## Backend Rules

- Follow Clean Architecture boundaries:
- `delivery/http -> usecase -> domain repository interface -> infrastructure persistence`

- Current backend layout:
- `internal/domain/<module>`
- `internal/usecase/<module>`
- `internal/delivery/http/dto`
- `internal/delivery/http/handler`
- `internal/delivery/http/response`
- `internal/delivery/http/router`
- `internal/infrastructure/...`

- Repository interfaces live in the consumer-owned domain package.
- Example: `internal/domain/todo/repository.go`

- Usecase files live by module.
- Example: `internal/usecase/todo/usecase.go`

- DTO naming must follow:
- `*_request.go`
- `*_response.go`

- Response helpers are split by responsibility:
- `success.go`
- `error.go`
- `error_mapper.go`

- Routes are registered through `RouteRegistrar` implementations and mounted under `/api/v1`.
- Do not register feature routes directly on the root engine.
- Health is also served under `/api/v1/health`.

- Handler rules:
- Parse HTTP input
- Call usecase
- Map DTO/response
- Do not place business rules in handlers

- Persistence rules:
- Infrastructure models must stay separate from domain entities.
- Mapping between persistence model and domain entity should be explicit.

- Test naming should stay module-qualified where useful:
- `internal/usecase/health/health_usecase_test.go`
- `internal/usecase/todo/todo_usecase_test.go`

## Frontend Rules

- Frontend is feature-based.
- Put feature-local code inside `apps/frontend/features/<feature>`.

- Current feature naming convention:
- `components/card.tsx`
- `components/form.tsx`
- `components/list.tsx`
- `hooks/use.ts`
- `types.ts`

- Keep shared code outside features only when it is genuinely shared:
- `components/`
- `hooks/`
- `services/`
- `types/api.ts`
- `constants/`
- `config/`

- Feature-specific types must stay in the feature folder.
- Generic API payload types can stay in `apps/frontend/types/api.ts`.

- Route constants live in `apps/frontend/constants/routes.ts`.
- Keep frontend API paths aligned with backend `/api/v1/*` routes.

## Before Editing

- Check the existing module pattern first.
- Match current naming before adding new files.
- If you add a backend module, follow the same domain/usecase/delivery structure already used by `todo` and `health`.
- If you add a frontend feature, follow the same `components/`, `hooks/use.ts`, and `types.ts` pattern already used by `health` and `todos`.
