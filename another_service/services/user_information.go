package services

import (
	"another_service/models"
	"strings"
)

func (s *Service) UserInformation(passportNumber string) (models.UserInformation, error) {
	passportNumber = strings.ReplaceAll(passportNumber, " ", "")
	passportNumber = strings.ReplaceAll(passportNumber, "-", "")
	passportNumber = strings.ReplaceAll(passportNumber, "+", "")
	return s.Repository.UserInformation(passportNumber)
}
