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
	GetAll() ([]models.TopUpOutput, error)
	GetByPenggunaID(IdPengguna string) ([]models.TopUpOutput, error)
	UpdateSaldoByID(IDPengguna string, ID string, amount int) (int, error)
	GetByID(ID string) (models.TopUpOutput, error)
	GetBySaldoByID(ID string) (models.SaldoPengguna, error)
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

	noTelp, _ := s.topUpRepository.GetTelponById(req.ID_Pengguna)

	req.KodeYukPay = "6969" + noTelp.NoTelpon

	req.TanggalTopUp = time.Now()

	err := s.topUpRepository.Create(req)

	return req.ID, req.KodeYukPay, err
}

func (s *topUpService) GetAll() ([]models.TopUpOutput, error) {
	topups, err := s.topUpRepository.GetAll()
	return topups, err
}

func (s *topUpService) GetByPenggunaID(IdPengguna string) ([]models.TopUpOutput, error) {
	topups, err := s.topUpRepository.GetByPenggunaID(IdPengguna)
	return topups, err
}

func (s *topUpService) GetByID(ID string) (models.TopUpOutput, error) {
	topups, err := s.topUpRepository.GetByID(ID)
	return topups, err
}

func (s *topUpService) GetBySaldoByID(ID string) (models.SaldoPengguna, error) {
	saldo, err := s.topUpRepository.GetSaldoByID(ID)
	return saldo, err
}

func (s *topUpService) UpdateSaldoByID(IDPengguna string, ID string, amount int) (int, error) {
	beforeAmount, _ := s.topUpRepository.GetSaldoByID(IDPengguna)
	updateAmount := beforeAmount.SaldoYukPay + amount
	err := s.topUpRepository.UpdateSaldoByID(IDPengguna, ID, updateAmount)
	return updateAmount, err
}
