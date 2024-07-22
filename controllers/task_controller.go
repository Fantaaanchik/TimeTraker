package controllers

import (
	"github.com/gin-gonic/gin"
	"time_traker/services"
)

type TaskController struct {
	Engine      *gin.Engine
	taskService *services.TaskService
}

func NewTaskController(service *services.TaskService, engine *gin.Engine) *TaskController {
	return &TaskController{
		taskService: service,
		Engine:      engine,
	}
}
