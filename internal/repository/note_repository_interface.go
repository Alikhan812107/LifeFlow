package repository

import (
	"Assignment3/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteRepository interface {
	Create(note models.Note) (models.Note, error)
	GetAll() ([]models.Note, error)
	GetByID(id primitive.ObjectID) (models.Note, error)
	Update(id primitive.ObjectID, note models.Note) (models.Note, error)
	Delete(id primitive.ObjectID) error
}