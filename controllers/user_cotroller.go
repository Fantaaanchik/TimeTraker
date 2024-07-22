package controllers

import (
	"github.com/gin-gonic/gin"
	"time_traker/services"
)

type UserController struct {
	Engine      *gin.Engine
	UserService *services.UserService
}

func NewUserController(services *services.UserService, engine *gin.Engine) *UserController {
	return &UserController{
		UserService: services,
		Engine:      engine,
	}
}
