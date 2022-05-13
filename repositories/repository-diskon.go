package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type DiskonRepository interface {
	GetAll() ([]models.Diskon, error)
}

type diskonRepository struct {
	db *gorm.DB
}

func NewDiskonRepository(db *gorm.DB) *diskonRepository {
	return &diskonRepository{db}
}

func (r *diskonRepository) GetAll() ([]models.Diskon, error) {
	var diskons []models.Diskon

	err := r.db.Table("diskon").Find(&diskons).Error
	return diskons, err
}
