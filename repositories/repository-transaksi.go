package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	GetAll() ([]models.Transaksi, error)
	GetTransaksiById(ID string) ([]models.Transaksi, error)
	GetTotalHarga() ([]models.Transaksi, error)
}

type transaksiRepository struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) *transaksiRepository {
	return &transaksiRepository{db}
}

func (r *transaksiRepository) GetAll() ([]models.Transaksi, error) {
	var transactions []models.Transaksi

	err := r.db.Table("transaksi").Select("transaksi.id_transaksi, transaksi.id_pengguna, pengguna.nama, varian.jenis, transaksi.totalHarga, transaksi.tanggal, transaksi.status").Joins("join varian on transaksi.id_varian=varian.id").Joins("join pengguna on transaksi.id_pengguna=pengguna.id").Scan(&transactions).Error
	return transactions, err
}

func (r *transaksiRepository) GetTransaksiById(ID string) ([]models.Transaksi, error) {
	var transactions []models.Transaksi

	err := r.db.Table("transaksi").Select("transaksi.id_transaksi, transaksi.id_pengguna, pengguna.nama, varian.jenis, transaksi.totalHarga, transaksi.tanggal, transaksi.status").Joins("join varian on transaksi.id_varian=varian.id").Joins("join pengguna on transaksi.id_pengguna=pengguna.id").Where("id_pengguna = ?", ID).Scan(&transactions).Error
	return transactions, err
}

func (r *transaksiRepository) GetTotalHarga() ([]models.Transaksi, error) {
	var transactions []models.Transaksi

	err := r.db.Table("transaksi").Select("totalHarga").Find(&transactions).Error

	return transactions, err
}
