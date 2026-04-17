package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/your-org/fullstack-template/apps/backend/internal/domain/todo"
)

type mappedError struct {
	err    error
	status int
}

type ErrorMapper struct {
	mappings []mappedError
}

var errorMappings = []mappedError{
	{
		err:    todo.ErrInvalidTitle,
		status: http.StatusBadRequest,
	},
}

var DefaultErrorMapper = NewErrorMapper(errorMappings...)

func NewErrorMapper(mappings ...mappedError) *ErrorMapper {
	return &ErrorMapper{mappings: mappings}
}

func WriteMappedError(c *gin.Context, err error, fallbackStatus int, fallbackMessage string) {
	if status, message, ok := DefaultErrorMapper.Map(err); ok {
		WriteError(c, status, message)
		return
	}

	WriteError(c, fallbackStatus, fallbackMessage)
}

func (m *ErrorMapper) Map(err error) (int, string, bool) {
	if err == nil {
		return 0, "", false
	}

	for _, mapping := range m.mappings {
		if errors.Is(err, mapping.err) {
			return mapping.status, err.Error(), true
		}
	}

	return 0, "", false
}
