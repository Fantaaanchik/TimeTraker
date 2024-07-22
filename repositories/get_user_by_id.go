package repositories

import "time_traker/models"

func (r *UserRepository) GetUserByID(id int) (models.User, error) {
	var user models.User

	if err := r.db.First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
