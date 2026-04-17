package repository

import (
	"context"

	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
)

type TodoRepository interface {
	List(ctx context.Context) ([]todo.Todo, error)
	Create(ctx context.Context, item todo.Todo) (todo.Todo, error)
}
