package services

import "time_traker/models"

func (s *UserService) UpdateUser(user models.User, id int) error {
	return s.userRepository.UpdateUser(user, id)
}
