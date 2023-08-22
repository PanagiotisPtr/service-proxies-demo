package middleware

import "github.com/gin-gonic/gin"

func WithErrorHandling(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		c.JSON(500, gin.H{"errors": c.Errors})
	}
}
