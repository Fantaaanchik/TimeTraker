package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Get details of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]string
// @Router /get_all_users [get]
func (uc UserController) GetAllUsers(c *gin.Context) {
	filter := c.Query("filter")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	users, err := uc.UserService.GetAllUsers(filter, page, pageSize)
	if err != nil {
		log.Println("cannot get all users, error description:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
