package services

func (s *TaskService) StopTask(id uint) error {
	return s.taskRepository.StopTask(id)
}
