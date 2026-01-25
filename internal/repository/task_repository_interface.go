package repository

import "Assignment3/internal/models"

// TaskRepository defines how task data is stored
type TaskRepository interface {
	Create(task models.Task) models.Task
	GetAll() []models.Task
	GetByID(id int) (models.Task, bool)
	Update(id int, task models.Task) (models.Task, bool)
	Delete(id int) bool
}
