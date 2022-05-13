package services

import (
	"strconv"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type TagihanService interface {
	Create(req *models.Tagihan) (models.Tagihan, error)
	GetPendingStatus(IdKategori string, IdPengguna string) ([]models.Tagihan, error)
	UpdateStatus(IdTagihan string) (models.Tagihan, error)
}

type tagihanService struct {
	tagihanRepository repositories.TagihanRepository
}

func NewTagihanService(tagihanRepository repositories.TagihanRepository) *tagihanService {
	return &tagihanService{tagihanRepository}
}

func (s *tagihanService) Create(req *models.Tagihan) (models.Tagihan, error) {
	num := s.tagihanRepository.GetDataCount()

	result := strconv.Itoa(int(num) + 1)
	req.ID_Tagihan = "TG" + "-" + result

	tagihan, err := s.tagihanRepository.Create(req)
	return tagihan, err
}

func (s *tagihanService) GetPendingStatus(IdKategori string, IdPengguna string) ([]models.Tagihan, error) {
	tagihans, err := s.tagihanRepository.GetPendingStatus(IdKategori, IdPengguna)
	return tagihans, err
}

func (s *tagihanService) UpdateStatus(IdTagihan string) (models.Tagihan, error) {
	tagihan, err := s.tagihanRepository.UpdateStatus(IdTagihan)
	return tagihan, err
}
