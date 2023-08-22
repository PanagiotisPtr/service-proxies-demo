package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panagiotisptr/service-proxies-demo/app"
	"github.com/panagiotisptr/service-proxies-demo/models"
)

type TaskController struct {
	app *app.App
}

func ProvideTaskController(app *app.App, e *gin.Engine) *TaskController {
	ctr := &TaskController{
		app: app,
	}

	e.GET("/tasks", ctr.ListTasks)
	e.GET("/tasks/:id", ctr.GetTask)
	e.POST("/tasks", ctr.CreateTask)
	e.PUT("/tasks", ctr.UpdateTask)
	e.DELETE("/tasks/:id", ctr.DeleteTask)

	return ctr
}

func (ctr *TaskController) ListTasks(c *gin.Context) {
	tasks, err := ctr.app.ListTasks(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (ctr *TaskController) GetTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	var task *models.Task
	task, err = ctr.app.GetTask(c, id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ctr *TaskController) CreateTask(c *gin.Context) {
	var task *models.Task
	var err error
	if err = c.ShouldBindJSON(&task); err != nil {
		c.Error(err)
		return
	}

	err = ctr.app.CreateTask(c, task)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ctr *TaskController) UpdateTask(c *gin.Context) {
	var task *models.Task
	var err error
	if err = c.ShouldBindJSON(&task); err != nil {
		c.Error(err)
		return
	}

	err = ctr.app.UpdateTask(c, task)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, task)
}

func (ctr *TaskController) DeleteTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Error(err)
		return
	}

	err = ctr.app.DeleteTask(c, id)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
