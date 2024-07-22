package handlers

import (
	"another_service/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Engine  *gin.Engine
	Service *services.Service
}

func NewHandler(services *services.Service, engine *gin.Engine) *Handler {
	return &Handler{
		Service: services,
		Engine:  engine,
	}
}
