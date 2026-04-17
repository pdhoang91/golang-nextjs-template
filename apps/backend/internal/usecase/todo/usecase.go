package todo

import (
	"context"
	"strings"

	"github.com/google/uuid"

	domain "github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
)

type CreateTodoInput struct {
	Title       string
	Description string
}

type TodoUsecase interface {
	List(ctx context.Context) ([]domain.Todo, error)
	Create(ctx context.Context, input CreateTodoInput) (domain.Todo, error)
}

type todoUsecase struct {
	todoRepository domain.TodoRepository
}

func NewTodoUsecase(todoRepository domain.TodoRepository) TodoUsecase {
	return &todoUsecase{
		todoRepository: todoRepository,
	}
}

func (u *todoUsecase) List(ctx context.Context) ([]domain.Todo, error) {
	return u.todoRepository.List(ctx)
}

func (u *todoUsecase) Create(ctx context.Context, input CreateTodoInput) (domain.Todo, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return domain.Todo{}, domain.ErrInvalidTitle
	}

	entity := domain.Todo{
		ID:          uuid.New(),
		Title:       title,
		Description: strings.TrimSpace(input.Description),
		Completed:   false,
	}

	return u.todoRepository.Create(ctx, entity)
}
