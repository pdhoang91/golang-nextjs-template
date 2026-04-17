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

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
	httpdto "github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/dto"
	handler "github.com/your-org/fullstack-template/apps/backend/internal/delivery/http/handler"
	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
	usetodo "github.com/your-org/fullstack-template/apps/backend/internal/usecase/todo"
)

type stubTodoUsecase struct{}

func (s stubTodoUsecase) List(_ context.Context) ([]todo.Todo, error) {
	return []todo.Todo{
		{
			ID:          uuid.MustParse(constants.TestUUIDTodoListID),
			Title:       constants.TestTodoSeedTitle,
			Description: constants.TestTodoSeedDescription,
			Completed:   false,
		},
	}, nil
}

func (s stubTodoUsecase) Create(_ context.Context, input usetodo.CreateTodoInput) (todo.Todo, error) {
	return todo.Todo{
		ID:          uuid.MustParse(constants.TestUUIDTodoCreateID),
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}, nil
}

func TestTodoHandlerList(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	h := handler.NewTodoHandler(stubTodoUsecase{})
	apiV1 := r.Group(constants.APIV1Prefix)
	h.RegisterRoutes(apiV1)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, constants.APIV1Prefix+constants.TodosRoute, nil)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), constants.TestTodoSeedTitle)
}

func TestTodoHandlerCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	h := handler.NewTodoHandler(stubTodoUsecase{})
	apiV1 := r.Group(constants.APIV1Prefix)
	h.RegisterRoutes(apiV1)

	body, err := json.Marshal(httpdto.CreateTodoRequest{
		Title:       constants.TestTodoCreateTitle,
		Description: constants.TestTodoCreateDescription,
	})
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, constants.APIV1Prefix+constants.TodosRoute, bytes.NewBuffer(body))
	req.Header.Set(constants.HeaderContentType, constants.ContentTypeJSON)

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	require.Contains(t, w.Body.String(), constants.TestTodoCreateTitle)
}
