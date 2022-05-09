package walletTopUp

import (
	"math/rand"
	"strconv"
	"time"
)

type topupservice interface {
	Create(req *WalletTopUp) (string, string, error)
	GetAll() ([]WalletTopUp, error)
	GetByPenggunaID(IdPengguna string) ([]WalletTopUp, error)
	UpdateSaldoByID(ID string, amount int) (int, error)
}

type service struct {
	repository topuprepository
}

func NewService(repository topuprepository) *service {
	return &service{repository}
}

func (s *service) Create(req *WalletTopUp) (string, string, error) {
	num := s.repository.GetDataCount()

	result := strconv.Itoa(int(num) + 1)
	req.ID = "WTP" + "-" + result
	rand.Seed(time.Now().UTC().UnixNano())

	req.KodeYukPay = strconv.Itoa(rand.Int())

	err := s.repository.Create(req)

	return req.ID, req.KodeYukPay, err
}

func (s *service) GetAll() ([]WalletTopUp, error) {
	topups, err := s.repository.GetAll()
	return topups, err
}

func (s *service) GetByPenggunaID(IdPengguna string) ([]WalletTopUp, error) {
	topups, err := s.repository.GetByPenggunaID(IdPengguna)
	return topups, err
}

func (s *service) UpdateSaldoByID(ID string, amount int) (int, error) {
	beforeAmount, _ := s.repository.GetSaldoByID(ID)
	updateAmount := beforeAmount.SaldoYukPay + amount
	err := s.repository.UpdateSaldoByID(ID, updateAmount)
	return updateAmount, err
}
