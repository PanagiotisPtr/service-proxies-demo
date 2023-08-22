package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panagiotisptr/service-proxies-demo/app"
	"github.com/panagiotisptr/service-proxies-demo/controller"
	"github.com/panagiotisptr/service-proxies-demo/middleware"
	taskmemoryrepo "github.com/panagiotisptr/service-proxies-demo/repo/memory"
	"github.com/panagiotisptr/service-proxies-demo/service"
	"github.com/panagiotisptr/service-proxies-demo/util"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func ProvideGinEngine(logger *zap.Logger) *gin.Engine {
	e := gin.Default()

	e.Use(middleware.WithUserID)
	e.Use(middleware.WithRequestID)
	e.Use(middleware.WithErrorHandling)
	e.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		traceStack := util.GetTraceStack(c)

		logger.With(
			zap.String("RequestID", c.GetString("RequestID")),
			zap.String("UserID", c.GetString("UserID")),
			zap.Strings("TraceStack", traceStack),
			zap.Any("Error", err),
		).Error("Oh no! Anyway...")
	}))

	return e
}

func Bootstrap(
	lc fx.Lifecycle,
	shutdowner fx.Shutdowner,
	e *gin.Engine,
	_ *controller.TaskController, // bootstrap controller
) {
	s := &http.Server{Addr: ":8080", Handler: e}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := s.ListenAndServe()
				if err != nil {
					shutdowner.Shutdown()
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			ProvideGinEngine,
			zap.NewDevelopment,
			taskmemoryrepo.ProvideMemoryTaskRepository,
			service.ProvideTaskService,
			app.ProvideApp,
			controller.ProvideTaskController,
		),
		fx.Invoke(Bootstrap),
		fx.WithLogger(
			func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			},
		),
	)

	app.Run()
}
