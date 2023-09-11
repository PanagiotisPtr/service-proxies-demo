package iservice

import (
	"context"

	"github.com/panagiotisptr/service-proxies-demo/models"
)

type TaskService interface {
	List(ctx context.Context) ([]*models.Task, error)
	Get(ctx context.Context, id int64) (*models.Task, error)
	Create(ctx context.Context, task *models.Task) error
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id int64) error
}
