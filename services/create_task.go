package services

import (
	"time_traker/models"
)

func (s *TaskService) StartTask(task models.Task) error {
	return s.taskRepository.StartTask(task)
}
