package todo

import "context"

type TodoRepository interface {
	List(ctx context.Context) ([]Todo, error)
	Create(ctx context.Context, item Todo) (Todo, error)
}
