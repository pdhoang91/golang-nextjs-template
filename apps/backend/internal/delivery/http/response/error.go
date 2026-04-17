package response

import "github.com/gin-gonic/gin"

func WriteError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		ErrorKey: message,
	})
}
