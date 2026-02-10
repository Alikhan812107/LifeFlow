package repository

import "Assignment3/internal/models"

type ActivityRepository interface {
	Create(activity models.Activity) (models.Activity, error)
	GetAll(userID string) ([]models.Activity, error)
}
