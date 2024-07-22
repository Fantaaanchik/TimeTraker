package repositories

import "time_traker/models"

func (r *TaskRepository) UpdateTask(task models.Task, id int) error {
	if err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(task).Error; err != nil {
		return err
	}
	return nil
}
