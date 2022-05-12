package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type MitraRepository interface {
	Create(req *models.Mitra) error
	GetAll() ([]models.Mitra, error)
	GetNamaPerusahaan(nama string) (models.Mitra, error)
	GetCount() int64
}

type mitraRepository struct {
	db *gorm.DB
}

func NewMitraRepository(db *gorm.DB) *mitraRepository {
	return &mitraRepository{db}
}

func (r *mitraRepository) Create(req *models.Mitra) error {
	err := r.db.Table("mitra").Create(req).Error
	return err
}

func (r *mitraRepository) GetAll() ([]models.Mitra, error) {
	var mitras []models.Mitra

	err := r.db.Table("mitra").Find(&mitras).Error
	return mitras, err
}

func (r *mitraRepository) GetNamaPerusahaan(nama string) (models.Mitra, error) {
	var mitra models.Mitra

	err := r.db.Table("mitra").Where("namaPerusahaan = ?", nama).Find(&mitra).Error

	return mitra, err
}

func (r *mitraRepository) GetCount() int64 {
	var num int64
	r.db.Table("mitra").Count(&num)
	return num
}
