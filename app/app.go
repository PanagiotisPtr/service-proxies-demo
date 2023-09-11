package app

import "github.com/panagiotisptr/service-proxies-demo/service/iservice"

type App struct {
	TaskSvc iservice.TaskService
}

func ProvideApp(taskSvc iservice.TaskService) *App {
	return &App{
		TaskSvc: taskSvc,
	}
}
