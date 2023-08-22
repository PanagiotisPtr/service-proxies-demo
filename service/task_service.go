//go:generate go-gen-proxy --interface github.com/panagiotisptr/service-proxies-demo/service.TaskService --package github.com/panagiotisptr/service-proxies-demo/service/proxy --name TaskService --output proxy/task_service.go
package service

import (
	"context"

	"github.com/eiiches/go-gen-proxy/pkg/interceptor"
	"github.com/panagiotisptr/service-proxies-demo/interceptors"
	"github.com/panagiotisptr/service-proxies-demo/models"
	"github.com/panagiotisptr/service-proxies-demo/repo"
	"github.com/panagiotisptr/service-proxies-demo/service/proxy"
	"go.uber.org/zap"
)

type TaskService interface {
	List(ctx context.Context) ([]*models.Task, error)
	Get(ctx context.Context, id int64) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id int64) error
}

type taskService struct {
	repo   repo.TaskRepository
	logger *zap.Logger
}

func ProvideTaskService(repo repo.TaskRepository, logger *zap.Logger) TaskService {
	l := logger.With(
		zap.String("service", "task"),
	)
	return &proxy.TaskService{
		Handler: &interceptor.InterceptingInvocationHandler{
			Delegate: &taskService{
				repo:   repo,
				logger: l,
			},
			Interceptor: &interceptors.TracingInterceptor{
				StructName: "TaskService",
				Logger:     l,
			},
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
