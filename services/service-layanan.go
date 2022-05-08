package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type LayananService interface {
	Create(req *models.Layanan) (string, bool, error)
	GetAll() ([]models.Layanan, error)
	GetByID(ID string) (models.Layanan, error)
}

type layananService struct {
	layananRepository repositories.LayananRepository
}

func NewLayananService(repositoryLayanan repositories.LayananRepository) *layananService {
	return &layananService{repositoryLayanan}
}

func (s *layananService) Create(req *models.Layanan) (string, bool, error) {
	layanan, err := s.layananRepository.GetByID(req.ID)
	if layanan.ID == req.ID {
		return "", true, err
	}

	err = s.layananRepository.Create(req)
	return req.ID, false, err
}
