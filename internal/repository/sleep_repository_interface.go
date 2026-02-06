package repository

import "Assignment3/internal/models"

type SleepRepository interface {
	Create(sleep models.Sleep) (models.Sleep, error)
	GetAll() ([]models.Sleep, error)
}
