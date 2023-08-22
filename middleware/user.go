package middleware

import "github.com/gin-gonic/gin"

func WithUserID(c *gin.Context) {
	userId := c.GetHeader("UserID")
	if userId == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("UserID", userId)
}
