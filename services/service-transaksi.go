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
	//UpdateStatusTransaksi(req *models.UpdateRequestTransaksi) (models.UpdateTransaksi, error)
}

type transaksiService struct {
	transaksiRepository repositories.TransaksiRepository
}

func NewTransaksiService(transaksiRepository repositories.TransaksiRepository) *transaksiService {
	return &transaksiService{transaksiRepository}
}

func (r *transaksiService) Create(req *models.NewTransaksi) (models.NewTransaksi, error) {
	num := r.transaksiRepository.GetCount()

	result := strconv.Itoa(int(num) + 1)
	req.ID_Transaksi = result

	req.Status = "Paid"
	transaksi, err := r.transaksiRepository.Create(req)
	return transaksi, err
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

// func (r *transaksiService) UpdateStatusTransaksi(req *models.UpdateRequestTransaksi) (models.UpdateTransaksi, error) {
// 	transaction, err := r.transaksiRepository.UpdateStatusTransaksi(req)
// 	return transaction, err
// }
