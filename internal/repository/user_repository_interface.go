package repository

import "Assignment3/internal/models"

type UserRepository interface {
	GetByID(id string) (*models.User, error)
	UpdateAvatar(id string, avatar string) error
}
