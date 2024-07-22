package services

import (
	"time_traker/models"
)

func (s *UserService) CreateUser(userData models.User) error {
	return s.userRepository.CreateUser(userData)
}
