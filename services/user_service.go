package services

import "time_traker/repositories"

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}
