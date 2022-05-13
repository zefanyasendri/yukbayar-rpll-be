package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type DiskonService interface {
	GetAll() ([]models.Diskon, error)
}

type diskonService struct {
	diskonRepository repositories.DiskonRepository
}

func NewDiskonService(diskonRepository repositories.DiskonRepository) *diskonService {
	return &diskonService{diskonRepository}
}

func (s *diskonService) GetAll() ([]models.Diskon, error) {
	diskons, err := s.diskonRepository.GetAll()
	return diskons, err
}
