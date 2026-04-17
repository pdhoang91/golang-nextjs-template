package usecase

import (
	"context"
	"strings"

	"github.com/google/uuid"

	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
	"github.com/your-org/fullstack-template/apps/backend/internal/repository"
)

type CreateTodoInput struct {
	Title       string
	Description string
}

type TodoUsecase interface {
	List(ctx context.Context) ([]todo.Todo, error)
	Create(ctx context.Context, input CreateTodoInput) (todo.Todo, error)
}

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(todoRepository repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		todoRepository: todoRepository,
	}
}

func (u *todoUsecase) List(ctx context.Context) ([]todo.Todo, error) {
	return u.todoRepository.List(ctx)
}

func (u *todoUsecase) Create(ctx context.Context, input CreateTodoInput) (todo.Todo, error) {
	title := strings.TrimSpace(input.Title)
	if title == "" {
		return todo.Todo{}, todo.ErrInvalidTitle
	}

	entity := todo.Todo{
		ID:          uuid.New(),
		Title:       title,
		Description: strings.TrimSpace(input.Description),
		Completed:   false,
	}

	return u.todoRepository.Create(ctx, entity)
}
