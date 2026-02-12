package service

import (
	"Assignment3/internal/models"
	"Assignment3/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) Create(user models.User) (*models.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) UpdateAvatar(id string, avatar string) error {
	return s.repo.UpdateAvatar(id, avatar)
}
