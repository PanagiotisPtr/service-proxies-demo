//go:generate proxygen --interface github.com/panagiotisptr/service-proxies-demo/service/iservice.TaskService --package proxy --name TaskService --output proxy/task_service.go
package service

import (
	"context"

	"github.com/panagiotisptr/proxygen/interceptor"
	"github.com/panagiotisptr/service-proxies-demo/interceptors"
	"github.com/panagiotisptr/service-proxies-demo/models"
	"github.com/panagiotisptr/service-proxies-demo/repo"
	"github.com/panagiotisptr/service-proxies-demo/service/iservice"
	"github.com/panagiotisptr/service-proxies-demo/service/proxy"
	"go.uber.org/zap"
)

type taskService struct {
	repo   repo.TaskRepository
	logger *zap.Logger
}

func ProvideTaskService(repo repo.TaskRepository, logger *zap.Logger) iservice.TaskService {
	l := logger.With(
		zap.String("service", "task"),
	)
	return &proxy.TaskService{
		Implementation: &taskService{
			repo:   repo,
			logger: l,
		},
		Interceptors: interceptor.InterceptorChain{
			interceptors.TracingInterceptor(logger, "TaskService"),
		},
	}
}

func (s *taskService) List(ctx context.Context) ([]*models.Task, error) {
	return s.repo.List(ctx)
}

func (s *taskService) Get(ctx context.Context, id int64) (*models.Task, error) {
	return s.repo.Get(ctx, id)
}

func (s *taskService) Create(ctx context.Context, task *models.Task) error {
	return s.repo.Create(ctx, task)
}

func (s *taskService) Update(ctx context.Context, task *models.Task) error {
	return s.repo.Update(ctx, task)
}

func (s *taskService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
