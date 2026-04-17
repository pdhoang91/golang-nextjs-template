package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
)

type stubTodoRepository struct {
	listResult   []todo.Todo
	createResult todo.Todo
	createInput  todo.Todo
}

func (s *stubTodoRepository) List(_ context.Context) ([]todo.Todo, error) {
	return s.listResult, nil
}

func (s *stubTodoRepository) Create(_ context.Context, item todo.Todo) (todo.Todo, error) {
	s.createInput = item
	if s.createResult.ID == uuid.Nil {
		s.createResult = item
	}
	return s.createResult, nil
}

func TestTodoUsecaseCreate(t *testing.T) {
	repo := &stubTodoRepository{}
	uc := NewTodoUsecase(repo)

	created, err := uc.Create(context.Background(), CreateTodoInput{
		Title:       "  First todo  ",
		Description: "  clean architecture  ",
	})

	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, created.ID)
	require.Equal(t, "First todo", created.Title)
	require.Equal(t, "clean architecture", created.Description)
	require.False(t, created.Completed)
}

func TestTodoUsecaseCreateEmptyTitle(t *testing.T) {
	repo := &stubTodoRepository{}
	uc := NewTodoUsecase(repo)

	_, err := uc.Create(context.Background(), CreateTodoInput{Title: "   "})

	require.ErrorIs(t, err, todo.ErrInvalidTitle)
}
