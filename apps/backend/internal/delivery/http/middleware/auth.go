package middleware

import "github.com/gin-gonic/gin"

// AuthPlaceholder là middleware mẫu để team thay bằng JWT/API key/... ở phiên bản sau.
func AuthPlaceholder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
