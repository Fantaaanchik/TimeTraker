package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	"log"
	"time_traker/config"
	"time_traker/controllers"
	"time_traker/db"
	_ "time_traker/docs"
	"time_traker/repositories"
	"time_traker/services"
)

// @title Time Tracker API
// @version 1.0
// @description This is a time tracking application.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file, %v", err.Error())
	}

	dbSettings := db.InitDB()

	r := gin.Default()

	userRepository := repositories.NewUserRepository(dbSettings)
	taskRepository := repositories.NewTaskRepository(dbSettings)

	userService := services.NewUserService(userRepository)
	taskService := services.NewTaskService(taskRepository)

	userController := controllers.NewUserController(userService, r)
	taskController := controllers.NewTaskController(taskService, r)

	userController.AllUserRoutes()
	taskController.AllTaskRoutes()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "page not found"})
	})

	err = r.Run(":8000")
	if err != nil {
		log.Fatal("fatal to run route, Error:", err.Error())
	}
}
