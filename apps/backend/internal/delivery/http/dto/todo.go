package dto

import (
	"time"

	"github.com/google/uuid"

	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
)

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type TodoResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ToTodoResponse(entity todo.Todo) TodoResponse {
	return TodoResponse{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Completed:   entity.Completed,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ToTodoResponses(items []todo.Todo) []TodoResponse {
	responses := make([]TodoResponse, 0, len(items))
	for _, item := range items {
		responses = append(responses, ToTodoResponse(item))
	}
	return responses
}
