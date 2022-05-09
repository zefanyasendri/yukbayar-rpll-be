package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type VarianService interface {
	GetAll() ([]models.Varian, error)
	GetByID(ID string) (models.Varian, error)
}

type varianService struct {
	varianRepository repositories.VarianRepository
}

func NewVarianService(varianRepository repositories.VarianRepository) *varianService {
	return &varianService{varianRepository}
}

func (s *varianService) GetAll() ([]models.Varian, error) {
	varians, err := s.varianRepository.GetAll()
	return varians, err
}

func (s *varianService) GetByID(ID string) (models.Varian, error) {
	varian, err := s.varianRepository.GetByID(ID)
	return varian, err
}
