package service

import (
	"Assignment3/internal/models"
	"Assignment3/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) Create(task models.Task) models.Task {
	return s.repo.Create(task)
}

func (s *TaskService) GetAll() []models.Task {
	return s.repo.GetAll()
}

func (s *TaskService) GetByID(id int) (models.Task, bool) {
	return s.repo.GetByID(id)
}

func (s *TaskService) Update(id int, task models.Task) (models.Task, bool) {
	return s.repo.Update(id, task)
}

func (s *TaskService) Delete(id int) bool {
	return s.repo.Delete(id)
}
