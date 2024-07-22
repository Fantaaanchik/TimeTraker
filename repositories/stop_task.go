package repositories

import (
	"time"
	"time_traker/models"
)

func (r *TaskRepository) StopTask(taskID uint) error {
	var task models.Task

	if err := r.db.First(&task, taskID).Error; err != nil {
		return err
	}

	if task.Completed {
		return nil
	}

	task.Completed = true
	duration := time.Since(task.StartedAt)
	task.Hours = int(duration.Hours())
	task.Minutes = int(duration.Minutes()) % 60

	return r.db.Save(&task).Error
}
