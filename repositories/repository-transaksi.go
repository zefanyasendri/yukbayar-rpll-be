package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	Create(req *models.NewTransaksi) (models.NewTransaksi, error)
	GetAll() ([]models.Transaksi, error)
	GetCount() int64
	GetTotalHarga() ([]models.Transaksi, error)
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

	err := r.db.Table("transaksi").Select("transaksi.id_transaksi, transaksi.id_pengguna, pengguna.nama, varian.jenis, transaksi.totalHarga, transaksi.tanggal, transaksi.status").Joins("join varian on transaksi.id_varian=varian.id").Joins("join pengguna on transaksi.id_pengguna=pengguna.id").Scan(&transactions).Error
	return transactions, err
}

func (r *transaksiRepository) GetTotalHarga() ([]models.Transaksi, error) {
	var transactions []models.Transaksi

	err := r.db.Table("transaksi").Select("totalHarga").Find(&transactions).Error

	return transactions, err
}
