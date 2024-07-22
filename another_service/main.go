package main

import (
	"another_service/db"
	"another_service/handlers"
	"another_service/repositories"
	"another_service/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	dbSettings := db.InitDB()

	r := gin.Default()

	userRepository := repositories.NewRepository(dbSettings)

	userService := services.NewService(userRepository)

	userController := handlers.NewHandler(userService, r)

	userController.AllRoutes()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "page not found"})
	})

	err := r.Run(":8081")
	if err != nil {
		log.Fatal("fatal to run route, Error:", err.Error())
	}
}
