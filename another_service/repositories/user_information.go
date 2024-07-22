package repositories

import (
	"another_service/models"
	"fmt"
)

func (r *Repository) UserInformation(passportNumber string) (models.UserInformation, error) {

	var userData models.UserInformation

	if err := r.db.Where("passport_number =?", passportNumber).First(&userData).Error; err != nil {
		wrappedError := fmt.Errorf("cannot get user data, error description: %v", err.Error())
		return models.UserInformation{}, wrappedError
	}

	return userData, nil
}
