package response

import "github.com/gin-gonic/gin"

func WriteSuccess(c *gin.Context, status int, data any) {
	c.JSON(status, gin.H{
		DataKey: data,
	})
}
