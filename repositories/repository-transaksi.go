package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	Create(req *models.NewTransaksi) (models.NewTransaksi, error)
	GetAll() ([]models.Transaksi, error)
	GetCount() int64
	GetTransaksiById(ID string) ([]models.Transaksi, error)
	GetTotalHarga() ([]models.Transaksi, error)
	//UpdateStatusTransaksi(req *models.UpdateRequestTransaksi) (models.UpdateTransaksi, error)
}

type transaksiRepository struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *transaksiRepository {
	return &transaksiRepository{db}
}

func (r *transaksiRepository) Create(req *models.NewTransaksi) (models.NewTransaksi, error) {
	err := r.db.Table("transaksi").Create(req).Error

	var transaksi models.NewTransaksi
	err = r.db.Table("transaksi").Where("id_transaksi = ?", req.ID_Transaksi).Scan(&transaksi).Error
	return transaksi, err
}

func (r *transaksiRepository) GetCount() int64 {
	var num int64

	r.db.Table("transaksi").Count(&num)
	return num
}

func (r *transaksiRepository) GetAll() ([]models.Transaksi, error) {
	var transactions []models.Transaksi

	err := r.db.Table("transaksi").Select("transaksi.id_transaksi, transaksi.id_pengguna, pengguna.nama, varian.jenis, transaksi.totalHarga, transaksi.tanggal, transaksi.status").Joins("join varian on transaksi.id_varian=varian.id").Joins("join pengguna on transaksi.id_pengguna=pengguna.id").Order("transaksi.tanggal DESC").Scan(&transactions).Error
	return transactions, err
}

func (r *transaksiRepository) GetTransaksiById(ID string) ([]models.Transaksi, error) {
	var transactions []models.Transaksi

	err := r.db.Table("transaksi").Select("transaksi.id_transaksi, transaksi.id_pengguna, pengguna.nama, varian.jenis, transaksi.totalHarga, transaksi.tanggal, transaksi.status").Joins("join varian on transaksi.id_varian=varian.id").Joins("join pengguna on transaksi.id_pengguna=pengguna.id").Where("id_pengguna = ?", ID).Order("transaksi.tanggal DESC").Scan(&transactions).Error
	return transactions, err
}

func (r *transaksiRepository) GetTotalHarga() ([]models.Transaksi, error) {
	var transactions []models.Transaksi

	err := r.db.Table("transaksi").Select("totalHarga").Find(&transactions).Error

	return transactions, err
}

// func (r *transaksiRepository) UpdateStatusTransaksi(req *models.UpdateRequestTransaksi) (models.UpdateTransaksi, error) {
// 	err := r.db.Table("transaksi").Where("id_pengguna = ?", req.ID_Pengguna).Where("id_varian", req.ID_varian).Update("status", "Paid").Error

// 	var transaction models.UpdateTransaksi
// 	err = r.db.Table("transaksi").Select("transaksi.id_transaksi, transaksi.id_pengguna, transaksi.id_varian, transaksi.status").Joins("join varian on transaksi.id_varian=varian.id").Joins("join pengguna on transaksi.id_pengguna=pengguna.id").Where("id_pengguna = ?", req.ID_Pengguna).Where("id_varian = ?", req.ID_varian).Find(&transaction).Error

// 	return transaction, err
// }
