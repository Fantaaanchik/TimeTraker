package services

func (s *TaskService) DeleteTask(id int) error {
	return s.taskRepository.DeleteTask(id)
}
