package services

import (
	"math/rand"
	"strconv"
	"time"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type TopUpService interface {
	Create(req *models.TopUp) (string, string, error)
	GetAll() ([]models.TopUp, error)
	GetByPenggunaID(IdPengguna string) ([]models.TopUp, error)
	UpdateSaldoByID(ID string, amount int) (int, error)
	GetByID(ID string) ([]models.TopUp, error)
}

type topUpService struct {
	topUpRepository repositories.TopUpRepository
}

func NewTopUpService(topUpRepository repositories.TopUpRepository) *topUpService {
	return &topUpService{topUpRepository}
}

func (s *topUpService) Create(req *models.TopUp) (string, string, error) {
	num := s.topUpRepository.GetDataCount()

	result := strconv.Itoa(int(num) + 1)
	req.ID = "WTP" + "-" + result
	rand.Seed(time.Now().UTC().UnixNano())

	req.KodeYukPay = strconv.Itoa(rand.Int())

	req.TanggalTopUp = time.Now()

	err := s.topUpRepository.Create(req)

	return req.ID, req.KodeYukPay, err
}

func (s *topUpService) GetAll() ([]models.TopUp, error) {
	topups, err := s.topUpRepository.GetAll()
	return topups, err
}

func (s *topUpService) GetByPenggunaID(IdPengguna string) ([]models.TopUp, error) {
	topups, err := s.topUpRepository.GetByPenggunaID(IdPengguna)
	return topups, err
}

func (s *topUpService) GetByID(ID string) ([]models.TopUp, error) {
	topups, err := s.topUpRepository.GetByID(ID)
	return topups, err
}

func (s *topUpService) UpdateSaldoByID(ID string, amount int) (int, error) {
	beforeAmount, _ := s.topUpRepository.GetSaldoByID(ID)
	updateAmount := beforeAmount.SaldoYukPay + amount
	err := s.topUpRepository.UpdateSaldoByID(ID, updateAmount)
	return updateAmount, err
}
