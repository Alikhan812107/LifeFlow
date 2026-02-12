package repository

import "Assignment3/internal/models"

type UserRepository interface {
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(user models.User) (*models.User, error)
	UpdateAvatar(id string, avatar string) error
	UpdateRole(id string, role string) error
}
