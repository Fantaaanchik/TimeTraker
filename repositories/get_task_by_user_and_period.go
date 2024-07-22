package repositories

import (
	"time"
	"time_traker/models"
)

func (r *TaskRepository) GetTasksByUserAndPeriod(userID uint, startDate, endDate time.Time) ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, startDate, endDate).
		Order("hours DESC, minutes DESC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
