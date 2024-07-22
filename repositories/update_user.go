package repositories

import "time_traker/models"

func (r *UserRepository) UpdateUser(user models.User, id int) error {

	if err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}

	return nil
}
