package services

import "another_service/repositories"

type Service struct {
	Repository *repositories.Repository
}

func NewService(repo *repositories.Repository) *Service {
	return &Service{Repository: repo}
}
