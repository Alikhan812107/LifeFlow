package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"Assignment3/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
}
