package util

import (
	"context"

	"github.com/gin-gonic/gin"
)

const TraceStackKey = "TraceStack"

func AddTraceToContext(ctx context.Context, trace string) context.Context {
	c, ok := ctx.(*gin.Context)
	if !ok {
		return ctx
	}

	stack, ok := c.Value(TraceStackKey).([]string)
	if !ok {
		stack = []string{}
	}
	stack = append(stack, trace)

	c.Set(TraceStackKey, stack)

	return c
}

func GetTraceStack(ctx context.Context) []string {
	stack, ok := ctx.Value(TraceStackKey).([]string)
	if !ok {
		stack = []string{}
	}

	return stack
}
