package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(constants.HeaderRequestID)
		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Set(constants.ContextRequestIDKey, requestID)
		c.Writer.Header().Set(constants.HeaderRequestID, requestID)
		c.Next()
	}
}
