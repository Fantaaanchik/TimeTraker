package services

import "time_traker/models"

func (s *TaskService) UpdateTask(task models.Task, id int) error {
	return s.taskRepository.UpdateTask(task, id)
}
