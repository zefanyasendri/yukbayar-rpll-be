package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type MitraService interface {
	Create(req *models.Mitra) (models.Mitra, error)
	//Create(req *models.Mitra) error
	GetAll() ([]models.Mitra, error)
}

type mitraService struct {
	repository repositories.MitraRepository
}

func NewMitraService(repository repositories.MitraRepository) *mitraService {
	return &mitraService{repository}
}

func (s *mitraService) Create(req *models.Mitra) (models.Mitra, error) {
	mitra, err := s.repository.Create(req)
	return mitra, err
}

func (s *mitraService) GetAll() ([]models.Mitra, error) {
	mitras, err := s.repository.GetAll()
	return mitras, err
}
