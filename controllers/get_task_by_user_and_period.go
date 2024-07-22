package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetWorkHours godoc
// @Summary Get tasks by user and period
// @Description Get tasks and their duration for a user within a specific period
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   user_id     query      string     true  "User ID"
// @Param   start_date  query      string     true  "Start date"
// @Param   end_date    query      string     true  "End date"
// @Success 200 {array} models.Task
// @Failure 500 {object} map[string]string
// @Router /work_hours_users [get]
func (tc TaskController) GetWorkHours(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	tasks, err := tc.taskService.GetTasksByUserAndPeriod(uint(userID), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
