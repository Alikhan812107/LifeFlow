package repository

import "Assignment3/internal/models"

type ActivityRepository interface {
	Create(activity models.Activity) (models.Activity, error)
	GetAll() ([]models.Activity, error)
}
