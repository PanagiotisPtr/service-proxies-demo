package memory

import (
	"context"

	"github.com/eiiches/go-gen-proxy/pkg/interceptor"
	"github.com/panagiotisptr/service-proxies-demo/interceptors"
	"github.com/panagiotisptr/service-proxies-demo/models"
	"github.com/panagiotisptr/service-proxies-demo/repo"
	"github.com/panagiotisptr/service-proxies-demo/repo/proxy"
	"go.uber.org/zap"
)

type MemoryTaskRepository struct {
	tasks  []*models.Task
	logger *zap.Logger
}

func ProvideMemoryTaskRepository(logger *zap.Logger) repo.TaskRepository {
	l := logger.With(
		zap.String("repository", "task"),
		zap.String("storage", "memory"),
	)
	return &proxy.TaskRepository{
		Handler: &interceptor.InterceptingInvocationHandler{
			Delegate: &MemoryTaskRepository{
				tasks:  []*models.Task{},
				logger: l,
			},
			Interceptor: &interceptors.TracingInterceptor{
				StructName: "MemoryTaskRepository",
				Logger:     l,
			},
		},
	}
}

func (r *MemoryTaskRepository) List(ctx context.Context) ([]*models.Task, error) {
	return r.tasks, nil
}

func (r *MemoryTaskRepository) Get(ctx context.Context, id int64) (*models.Task, error) {
	return r.tasks[id], nil // out of bounds error
}

func (r *MemoryTaskRepository) Create(ctx context.Context, task *models.Task) error {
	task.Id = int64(len(r.tasks))
	r.tasks = append(r.tasks, task)

	return nil
}

func (r *MemoryTaskRepository) Update(ctx context.Context, task *models.Task) error {
	// also out of bounds error - bad stuff
	r.tasks[task.Id] = task

	return nil
}

func (r *MemoryTaskRepository) Delete(ctx context.Context, id int64) error {
	for i, t := range r.tasks {
		if t.Id == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return nil
}
