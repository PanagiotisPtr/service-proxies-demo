package interceptors

import (
	"context"
	"fmt"

	"github.com/panagiotisptr/service-proxies-demo/util"
	"go.uber.org/zap"
)

type TracingInterceptor struct {
	Logger     *zap.Logger
	StructName string
}

func (ti *TracingInterceptor) Intercept(receiver interface{}, method string, args []interface{}, delegate func([]interface{}) []interface{}) (rets []interface{}) {
	logger := ti.Logger.With(
		zap.String("method", method),
		zap.String("struct", ti.StructName),
	)

	var ctx context.Context
	ctxIdx := -1
	for idx, arg := range args {
		if _, ok := arg.(context.Context); ok {
			ctx = arg.(context.Context)
			ctxIdx = idx
			break
		}
	}
	if ctx != nil {
		if userID, ok := ctx.Value("UserID").(string); ok {
			logger = logger.With(
				zap.String("UserID", userID),
			)
		} else {
			logger.Info("no user id")
		}

		if requestID, ok := ctx.Value("RequestID").(string); ok {
			logger = logger.With(
				zap.String("RequestID", requestID),
			)
		} else {
			logger.Info("no request id")
		}

		args[ctxIdx] = util.AddTraceToContext(
			ctx,
			fmt.Sprintf("%s.%s", ti.StructName, method),
		)
	}

	logger.Info("calling method")

	return delegate(args)
}
