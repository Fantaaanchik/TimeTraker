package services

import "time_traker/models"

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.userRepository.GetUserByID(id)
}
