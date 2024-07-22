package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time_traker/models"
)

// UpdateUser godoc
// @Summary Update user by ID
// @Description Update details of a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path      string     true  "User ID"
// @Param   user   body      models.User  true  "User info"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /update_user/{id} [put]
func (uc UserController) UpdateUser(c *gin.Context) {
	var user models.User

	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong type of id, need number"})
		return
	}
	err = c.ShouldBindJSON(&user)
	if err != nil {
		log.Println("cannot update user, error description: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err = uc.UserService.UpdateUser(user, id)
	if err != nil {
		log.Println("cannot update user, error description:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "User successfully updated!"})
}
