package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type LayananService interface {
	GetAll() ([]models.Layanan, error)
}

type layananService struct {
	layananRepository repositories.LayananRepository
}

func NewLayananService(repositoryLayanan repositories.LayananRepository) *layananService {
	return &layananService{repositoryLayanan}
}

func (s *layananService) GetAll() ([]models.Layanan, error) {
	varians, err := s.layananRepository.GetAll()
	return varians, err
}
