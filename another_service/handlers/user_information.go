package handlers

import (
	"another_service/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h Handler) UserInformation(c *gin.Context) {
	passportNumber := c.Query("passport_number")

	userData, err := h.Service.UserInformation(passportNumber)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	log.Println(userData)

	user := models.User{
		ID:             userData.ID,
		PassportNumber: userData.PassportNumber,
		Surname:        userData.Surname,
		Name:           userData.Name,
		Patronymic:     userData.Patronymic,
		Address:        userData.Address,
	}

	log.Println(user)
	c.JSON(http.StatusOK, user)
}
