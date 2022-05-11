package services

import (
	"strconv"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type MitraService interface {
	Create(req *models.Mitra) (models.Mitra, error)
	//Create(req *models.Mitra) error
	GetAll() ([]models.Mitra, error)
}

type mitraService struct {
	mitraRepository repositories.MitraRepository
}

func NewMitraService(mitraRepository repositories.MitraRepository) *mitraService {
	return &mitraService{mitraRepository}
}

func (s *mitraService) Create(req *models.Mitra) (models.Mitra, error) {
	num := s.mitraRepository.GetCount()

	result := strconv.Itoa(int(num) + 1)
	req.ID = "WTP" + "-" + result
	mitra, err := s.mitraRepository.Create(req)
	return mitra, err
}

func (s *mitraService) GetAll() ([]models.Mitra, error) {
	mitras, err := s.mitraRepository.GetAll()
	return mitras, err
}
