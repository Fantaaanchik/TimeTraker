package services

import "time_traker/models"

func (s *UserService) GetAllUsers(filter string, page, pageSize int) ([]models.User, error) {
	return s.userRepository.GetAllUsers(filter, page, pageSize)
}
