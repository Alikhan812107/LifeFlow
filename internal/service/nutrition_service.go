package service

import (
	"Assignment3/internal/models"
	"Assignment3/internal/repository"
)

type NutritionService struct {
	repo repository.NutritionRepository
}

func NewNutritionService(repo repository.NutritionRepository) *NutritionService {
	return &NutritionService{repo: repo}
}

func (s *NutritionService) Create(nutrition models.Nutrition) (models.Nutrition, error) {
	return s.repo.Create(nutrition)
}

func (s *NutritionService) GetAll() ([]models.Nutrition, error) {
	return s.repo.GetAll()
}
