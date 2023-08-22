package app

import (
	"context"

	"github.com/panagiotisptr/service-proxies-demo/models"
)

func (a *App) ListTasks(ctx context.Context) ([]*models.Task, error) {
	return a.TaskSvc.List(ctx)
}

func (a *App) GetTask(ctx context.Context, id int64) (*models.Task, error) {
	return a.TaskSvc.Get(ctx, id)
}

func (a *App) CreateTask(ctx context.Context, task *models.Task) error {
	return a.TaskSvc.Create(ctx, task)
}

func (a *App) UpdateTask(ctx context.Context, task *models.Task) error {
	return a.TaskSvc.Update(ctx, task)
}

func (a *App) DeleteTask(ctx context.Context, id int64) error {
	return a.TaskSvc.Delete(ctx, id)
}
