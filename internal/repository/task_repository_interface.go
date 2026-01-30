package repository

import (
	"Assignment3/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository interface {
	Create(task models.Task) (models.Task, error)
	GetAll() ([]models.Task, error)
	GetByID(id primitive.ObjectID) (models.Task, error)
	Update(id primitive.ObjectID, task models.Task) (models.Task, error)
	Delete(id primitive.ObjectID) error
}
