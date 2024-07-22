package repositories

import "time_traker/models"

func (r *UserRepository) CreateUser(user models.User) error {

	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
