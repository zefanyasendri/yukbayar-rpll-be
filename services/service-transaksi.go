package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type TransaksiService interface {
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

func (r *transaksiService) GetAll() ([]models.Transaksi, error) {
	transactions, err := r.transaksiRepository.GetAll()
	return transactions, err
}

func (r *transaksiService) GetTransaksiById(ID string) ([]models.Transaksi, error) {
	transactions, err := r.transaksiRepository.GetTransaksiById(ID)
	return transactions, err
}

func (r *transaksiService) GetTotalHarga() (int, error) {
	transactions, err := r.transaksiRepository.GetTotalHarga()

	var totalHarga int

	totalHarga = 0

	for i := 0; i < len(transactions); i++ {
		totalHarga += transactions[i].TotalHarga
	}

	return totalHarga, err
}
