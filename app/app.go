package app

import "github.com/panagiotisptr/service-proxies-demo/service"

type App struct {
	TaskSvc service.TaskService
}

func ProvideApp(taskSvc service.TaskService) *App {
	return &App{
		TaskSvc: taskSvc,
	}
}
