package controllers

import "github.com/gin-gonic/gin"

func (tc TaskController) AllTaskRoutes() {
	r := gin.Default()

	r.POST("/start_task", tc.StartTask)
	r.GET("/stop_task/:id", tc.StopTask)
	r.GET("/work_hours_users", tc.GetWorkHours)
}
