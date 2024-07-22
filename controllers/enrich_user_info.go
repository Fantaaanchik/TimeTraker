package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time_traker/models"
)

func (uc UserController) EnrichUserInfo(passportNumber string, c *gin.Context) (models.User, int, error) {
	var response struct {
		User models.User `json:"user"`
	}
	resp, err := http.Get(fmt.Sprintf("http://localhost:8081/info?passport_number=%s", passportNumber))
	if err != nil {
		return models.User{}, 500, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.User{}, resp.StatusCode, err
	}

	log.Println("Response Body: ", string(body))

	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.User{}, resp.StatusCode, err
	}
	log.Println("------------------------------------------------------>", response.User.ID)
	enrichedUser := response.User
	return enrichedUser, resp.StatusCode, nil
}
