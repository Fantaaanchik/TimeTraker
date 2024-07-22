package services

import (
	"time"
	"time_traker/models"
)

func (s *TaskService) GetTasksByUserAndPeriod(userID uint, startDate, endDate time.Time) ([]models.Task, error) {
	return s.taskRepository.GetTasksByUserAndPeriod(userID, startDate, endDate)
}
