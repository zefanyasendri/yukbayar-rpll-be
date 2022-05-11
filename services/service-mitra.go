package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type MitraService interface {
	Create(req *models.Mitra) (string, bool, error)
	GetAll() ([]models.Mitra, error)
}

type mitraService struct {
	repository repositories.MitraRepository
}

func NewMitraService(repository repositories.MitraRepository) *mitraService {
	return &mitraService{repository}
}

func (s *mitraService) Create(req *models.Mitra) (string, bool, error) {
	mitra, err := s.repository.GetByIdLayanan(req.ID_Layanan)

	if mitra.ID_Layanan == req.ID_Layanan {
		return "", true, err
	}

	err = s.repository.Create(req)
	return req.ID_Layanan, false, err
}

func (s *mitraService) GetAll() ([]models.Mitra, error) {
	mitras, err := s.repository.GetAll()
	return mitras, err
}
