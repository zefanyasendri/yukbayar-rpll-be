package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type KategoriLayananService interface {
	GetAll() ([]models.KategoriLayanan, error)
	GetByID(ID string) (models.KategoriLayanan, error)
	GetIDByKategori(ID string) ([]models.Varian, error)
}

type kategoriLayananService struct {
	kategoriLayananRepository repositories.KategoriLayananRepository
}

func NewKategoriLayananService(kategoriLayananRepository repositories.KategoriLayananRepository) *kategoriLayananService {
	return &kategoriLayananService{kategoriLayananRepository}
}

func (s *kategoriLayananService) GetAll() ([]models.KategoriLayanan, error) {
	varians, err := s.kategoriLayananRepository.GetAll()
	return varians, err
}

func (s *kategoriLayananService) GetByID(ID string) (models.KategoriLayanan, error) {
	kategoriLayanan, err := s.kategoriLayananRepository.GetByID(ID)
	return kategoriLayanan, err
}

func (s *kategoriLayananService) GetIDByKategori(ID string) ([]models.Varian, error) {
	var varians []models.Varian

	varians, err := s.kategoriLayananRepository.GetIDByKategori(ID)
	return varians, err
}
