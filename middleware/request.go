package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func WithRequestID(c *gin.Context) {
	c.Set("RequestID", uuid.New().String())
}
