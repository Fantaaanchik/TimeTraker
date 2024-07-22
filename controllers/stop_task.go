package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// StopTask godoc
// @Summary Stop a task for a user
// @Description Stop tracking time for a task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   id     path      string     true  "Task ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stop_task/{id} [post]
func (tc TaskController) StopTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	err = tc.taskService.StopTask(uint(taskID))
	if err != nil {
		log.Println("cannot stop task, error description: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}
	c.JSON(http.StatusOK, gin.H{"description": "task stopped successfully"})
}
