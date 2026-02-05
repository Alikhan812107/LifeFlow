package service

import (
	"Assignment3/internal/models"
	"Assignment3/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(note models.Note) (models.Note, error) {
	return s.repo.Create(note)
}

func (s *NoteService) GetAll() ([]models.Note, error) {
	return s.repo.GetAll()
}

func (s *NoteService) GetByID(id primitive.ObjectID) (models.Note, error) {
	return s.repo.GetByID(id)
}

func (s *NoteService) Update(id primitive.ObjectID, note models.Note) (models.Note, error) {
	return s.repo.Update(id, note)
}

func (s *NoteService) Delete(id primitive.ObjectID) error {
	return s.repo.Delete(id)
}