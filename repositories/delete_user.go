package repositories

import "time_traker/models"

func (r *UserRepository) DeleteUser(id int) error {

	if err := r.db.Delete(models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
