package repository

import (
	"Assignment3/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetByID(id primitive.ObjectID) (models.User, error)
}
