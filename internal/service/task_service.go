package service

import (
	"Assignment3/internal/models"
	"Assignment3/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(task models.Task) (models.Task, error) {
	return s.repo.Create(task)
}

func (s *TaskService) GetAll() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetByID(id primitive.ObjectID) (models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) Update(id primitive.ObjectID, task models.Task) (models.Task, error) {
	return s.repo.Update(id, task)
}

func (s *TaskService) Delete(id primitive.ObjectID) error {
	return s.repo.Delete(id)
}
