package repository

import "Assignment3/internal/models"

type NutritionRepository interface {
	Create(nutrition models.Nutrition) (models.Nutrition, error)
	GetAll() ([]models.Nutrition, error)
}
