package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type LayananRepository interface {
	GetAll() ([]models.Layanan, error)
}

type layananRepository struct {
	db *gorm.DB
}

func NewLayananRepository(db *gorm.DB) *layananRepository {
	return &layananRepository{db}
}

func (r *layananRepository) GetAll() ([]models.Layanan, error) {
	var layanans []models.Layanan

	err := r.db.Table("layanan").Find(&layanans).Error
	return layanans, err
}
