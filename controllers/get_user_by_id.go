package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get details of a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id   path      string     true  "User ID"
// @Success 200 {object} models.User
// @Failure 500 {object} map[string]string
// @Router /get_user_by_id/{id} [get]
func (uc UserController) GetUserByID(c *gin.Context) {
	idS := c.Param("id")

	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong type of id, need number"})
		return
	}

	user, err := uc.UserService.GetUserByID(id)
	if err != nil {
		log.Println("cannot get user by id, error description: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
