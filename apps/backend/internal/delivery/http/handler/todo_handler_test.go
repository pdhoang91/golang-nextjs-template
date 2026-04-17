package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	httpdto "github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/dto"
	handler "github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/handler"
	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
	"github.com/your-org/fullstack-template/apps/backend/internal/usecase"
)

type stubTodoUsecase struct{}

func (s stubTodoUsecase) List(_ context.Context) ([]todo.Todo, error) {
	return []todo.Todo{
		{
			ID:          uuid.MustParse("7c88a7f6-2e65-4c6f-b6df-8bfb84d70b5e"),
			Title:       "Seeded todo",
			Description: "Returned from stub usecase",
			Completed:   false,
		},
	}, nil
}

func (s stubTodoUsecase) Create(_ context.Context, input usecase.CreateTodoInput) (todo.Todo, error) {
	return todo.Todo{
		ID:          uuid.MustParse("4d9636a8-f035-4e72-9dd6-6d26da6d5ad3"),
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}, nil
}

func TestTodoHandlerList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	h := handler.NewTodoHandler(stubTodoUsecase{})
	r.GET("/api/v1/todos", h.List)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), "Seeded todo")
}

func TestTodoHandlerCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	h := handler.NewTodoHandler(stubTodoUsecase{})
	r.POST("/api/v1/todos", h.Create)

	body, err := json.Marshal(httpdto.CreateTodoRequest{
		Title:       "Write docs",
		Description: "Document the template",
	})
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	require.Contains(t, w.Body.String(), "Write docs")
}
