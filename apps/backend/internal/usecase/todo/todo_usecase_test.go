package todo

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
	domain "github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
)

type stubTodoRepository struct {
	listResult   []domain.Todo
	createResult domain.Todo
	createInput  domain.Todo
}

func (s *stubTodoRepository) List(_ context.Context) ([]domain.Todo, error) {
	return s.listResult, nil
}

func (s *stubTodoRepository) Create(_ context.Context, item domain.Todo) (domain.Todo, error) {
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
		Title:       constants.TestTodoFirstTitle,
		Description: constants.TestTodoFirstDescription,
	})

	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, created.ID)
	require.Equal(t, constants.TestTodoTrimmedTitle, created.Title)
	require.Equal(t, constants.TestTodoTrimmedDescription, created.Description)
	require.False(t, created.Completed)
}

func TestTodoUsecaseCreateEmptyTitle(t *testing.T) {
	repo := &stubTodoRepository{}
	uc := NewTodoUsecase(repo)

	_, err := uc.Create(context.Background(), CreateTodoInput{Title: "   "})

	require.ErrorIs(t, err, domain.ErrInvalidTitle)
}
