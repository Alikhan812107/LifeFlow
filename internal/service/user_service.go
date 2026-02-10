package service

import (
	"Assignment3/internal/models"
	"Assignment3/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(user models.User) (models.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) GetByEmail(email string) (models.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetByID(id primitive.ObjectID) (models.User, error) {
	return s.repo.GetByID(id)
}
