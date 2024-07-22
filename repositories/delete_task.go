package repositories

import "time_traker/models"

func (r *TaskRepository) DeleteTask(id int) error {

	if err := r.db.Delete(models.Task{}, id).Error; err != nil {
		return err
	}

	return nil
}
