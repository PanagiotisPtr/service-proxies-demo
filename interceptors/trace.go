package interceptors

import (
	"context"
	"fmt"

	"github.com/panagiotisptr/proxygen/interceptor"
	"github.com/panagiotisptr/service-proxies-demo/util"
	"go.uber.org/zap"
)

func TracingInterceptor(
	logger *zap.Logger,
	structName string,
) interceptor.Interceptor {
	return func(method string, next interceptor.Handler) interceptor.Handler {
		logger := logger.With(
			zap.String("method", method),
			zap.String("struct", structName),
		)

		return func(args []interface{}) []interface{} {
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
					fmt.Sprintf("%s.%s", structName, method),
				)
			}

			logger.Info("calling method")

			return next(args)
		}
	}
}
