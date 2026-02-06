package service

import (
	"Assignment3/internal/models"
	"Assignment3/internal/repository"
)

type ActivityService struct {
	repo repository.ActivityRepository
}

func NewActivityService(repo repository.ActivityRepository) *ActivityService {
	return &ActivityService{repo: repo}
}

func (s *ActivityService) Create(activity models.Activity) (models.Activity, error) {
	return s.repo.Create(activity)
}

func (s *ActivityService) GetAll() ([]models.Activity, error) {
	return s.repo.GetAll()
}
