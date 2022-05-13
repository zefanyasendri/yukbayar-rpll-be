package services

import (
	"strconv"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type TransaksiService interface {
	Create(req *models.NewTransaksi) (models.NewTransaksi, error)
	GetAll() ([]models.Transaksi, error)
	GetTransaksiById(ID string) ([]models.Transaksi, error)
	GetTotalHarga() (int, error)
}

type transaksiService struct {
	transaksiRepository repositories.TransaksiRepository
}

func NewTransaksiService(transaksiRepository repositories.TransaksiRepository) *transaksiService {
	return &transaksiService{transaksiRepository}
}

func (s *transaksiService) Create(req *models.NewTransaksi) (models.NewTransaksi, error) {
	num := s.transaksiRepository.GetCount()

	result := strconv.Itoa(int(num) + 1)
	req.ID_Transaksi = result

	req.Status = "Paid"
	transaksi, err := s.transaksiRepository.Create(req)
	return transaksi, err
}

func (s *transaksiService) GetAll() ([]models.Transaksi, error) {
	transactions, err := s.transaksiRepository.GetAll()
	return transactions, err
}

func (s *transaksiService) GetTransaksiById(ID string) ([]models.Transaksi, error) {
	transactions, err := s.transaksiRepository.GetTransaksiById(ID)
	return transactions, err
}

func (s *transaksiService) GetTotalHarga() (int, error) {
	transactions, err := s.transaksiRepository.GetTotalHarga()

	var totalHarga int

	totalHarga = 0

	for i := 0; i < len(transactions); i++ {
		totalHarga += transactions[i].TotalHarga
	}

	return totalHarga, err
}
