package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type MitraRepository interface {
	Create(req *models.Mitra) (models.Mitra, error)
	//Create(req *models.Mitra) error
	GetAll() ([]models.Mitra, error)
	// GetByIdLayanan(ID string) (models.Mitra, error)
}

type mitraRepository struct {
	db *gorm.DB
}

func NewMitraRepository(db *gorm.DB) *mitraRepository {
	return &mitraRepository{db}
}

func (r *mitraRepository) Create(req *models.Mitra) (models.Mitra, error) {
	err := r.db.Table("mitra").Create(req).Error
	var modelMitra models.Mitra
	r.db.Table("mitra").Where("id = ?", req.ID).Find(&modelMitra)
	return modelMitra, err
}

func (r *mitraRepository) GetAll() ([]models.Mitra, error) {
	var mitras []models.Mitra

	err := r.db.Table("mitra").Find(&mitras).Error
	return mitras, err
}

// func (r *mitraRepository) GetByIdLayanan(ID string) (models.Mitra, error) {
// 	var mitras models.Mitra

// 	err := r.db.Table("mitra").Where("id_layanan = ?", ID).Scan(&mitras).Error

// 	return mitras, err
// }
