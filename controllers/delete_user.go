package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// DeleteUser godoc
// @Summary Delete user by ID
// @Description Delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id   path      string     true  "User ID"
// @Success 204 {object} nil
// @Failure 500 {object} map[string]string
// @Router /delete_user/{id} [delete]
func (uc UserController) DeleteUser(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err = uc.UserService.DeleteUser(id)
	if err != nil {
		log.Println("cannot delete user, error description:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"description": "user deleted successfully!"})
}
