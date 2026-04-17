package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/your-org/fullstack-template/apps/backend/internal/constants"
)

func RequestLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startedAt := time.Now()
		c.Next()

		requestID, _ := c.Get(constants.ContextRequestIDKey)

		logger.Info(constants.LogHTTPRequest,
			constants.LogFieldRequestID, requestID,
			constants.LogFieldMethod, c.Request.Method,
			constants.LogFieldPath, c.Request.URL.Path,
			constants.LogFieldStatus, c.Writer.Status(),
			constants.LogFieldClientIP, c.ClientIP(),
			constants.LogFieldDurationMS, time.Since(startedAt).Milliseconds(),
		)
	}
}
