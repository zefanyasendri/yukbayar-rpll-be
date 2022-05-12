package services

import (
	"strconv"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type MitraService interface {
	Create(req *models.Mitra) (models.Mitra, bool, error)
	GetAll() ([]models.Mitra, error)
}

type mitraService struct {
	mitraRepository repositories.MitraRepository
}

func NewMitraService(mitraRepository repositories.MitraRepository) *mitraService {
	return &mitraService{mitraRepository}
}

func (s *mitraService) Create(req *models.Mitra) (models.Mitra, bool, error) {
	num := s.mitraRepository.GetCount()

	result := strconv.Itoa(int(num) + 1)
	req.ID = result
	mitra, err := s.mitraRepository.GetNamaPerusahaan(req.NamaPerusahaan)

	if mitra.NamaPerusahaan == req.NamaPerusahaan {
		return mitra, true, err
	}
	mitra, err = s.mitraRepository.Create(req)

	return mitra, false, err
}

func (s *mitraService) GetAll() ([]models.Mitra, error) {
	mitras, err := s.mitraRepository.GetAll()
	return mitras, err
}
