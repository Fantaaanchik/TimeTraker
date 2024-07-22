package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time_traker/models"
)

// StartTask godoc
// @Summary Start a task for a user
// @Description Start tracking time for a task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   task     body    models.Task     true        "Task info"
// @Success 201 {object} models.Task
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /start_task [post]
func (tc TaskController) StartTask(c *gin.Context) {
	var task models.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err = tc.taskService.StartTask(task)
	if err != nil {
		log.Println("cannot start task, error description: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "task created successfully"})
}
