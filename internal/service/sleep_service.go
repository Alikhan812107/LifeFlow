package service

import (
	"Assignment3/internal/models"
	"Assignment3/internal/repository"
)

type SleepService struct {
	repo repository.SleepRepository
}

func NewSleepService(repo repository.SleepRepository) *SleepService {
	return &SleepService{repo: repo}
}

func (s *SleepService) Create(sleep models.Sleep) (models.Sleep, error) {
	return s.repo.Create(sleep)
}

func (s *SleepService) GetAll() ([]models.Sleep, error) {
	return s.repo.GetAll()
}
