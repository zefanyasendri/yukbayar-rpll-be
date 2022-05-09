package services

import (
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/repositories"
)

type TransaksiService interface {
	GetAll() ([]models.Transaksi, error)
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
