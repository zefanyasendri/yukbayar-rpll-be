package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type LayananRepository interface {
	Create(req *models.Layanan) error
	GetAll() ([]models.Layanan, error)
	GetByID(ID string) (models.Layanan, error)
}

type layananRepository struct {
	db *gorm.DB
}

func NewLayananRepository(db *gorm.DB) *layananRepository {
	return &layananRepository{db}
}

func (r *layananRepository) Create(req *models.Layanan) error {
	err := r.db.Table("layanan").Create(req).Error
	return err
}

func (r *layananRepository) GetAll() ([]models.Layanan, error) {
	var layanans []models.Layanan

	err := r.db.Table("layanan").Find(&layanans).Error
	return layanans, err
}

func (r *layananRepository) GetByID(ID string) (models.Layanan, error) {
	var layanan models.Layanan

	err := r.db.Table("layanan").Where("id = ?", ID).Scan(&layanan).Error
	return layanan, err
}
