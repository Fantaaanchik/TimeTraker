package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time_traker/models"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.CreateUser true "User Info"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /create_user [post]
func (uc UserController) CreateUser(c *gin.Context) {
	var userInfo models.CreateUser

	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	userData, statusCode, err := uc.EnrichUserInfo(userInfo.PassportNumber, c)
	if err != nil {
		log.Println(err.Error())
		c.JSON(statusCode, gin.H{"error": "bad request"})
		return
	}

	err = uc.UserService.CreateUser(userData)
	if err != nil {
		log.Println("cannot create user, error description: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "Successfully create new user"})
}
