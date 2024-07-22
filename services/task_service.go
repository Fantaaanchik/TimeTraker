package services

import "time_traker/repositories"

type TaskService struct {
	taskRepository *repositories.TaskRepository
}

func NewTaskService(taskRepo *repositories.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepo,
	}
}
