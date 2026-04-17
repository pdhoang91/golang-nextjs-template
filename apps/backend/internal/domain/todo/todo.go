package todo

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidTitle = errors.New("todo title is required")

type Todo struct {
	ID          uuid.UUID
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
