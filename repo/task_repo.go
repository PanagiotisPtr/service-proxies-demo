//go:generate proxygen --interface github.com/panagiotisptr/service-proxies-demo/repo.TaskRepository --package proxy --name TaskRepository --output proxy/task_repo.go
package repo

import (
	"context"

	"github.com/panagiotisptr/service-proxies-demo/models"
)

type TaskRepository interface {
	List(ctx context.Context) ([]*models.Task, error)
	Get(ctx context.Context, id int64) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id int64) error
}
