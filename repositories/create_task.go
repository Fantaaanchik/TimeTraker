package repositories

import (
	"time"
	"time_traker/models"
)

func (r *TaskRepository) StartTask(task models.Task) error {
	task.StartedAt = time.Now()
	if err := r.db.Create(&task).Error; err != nil {
		return err
	}

	return nil
}
