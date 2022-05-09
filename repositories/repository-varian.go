package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type VarianRepository interface {
	GetAll() ([]models.Varian, error)
	GetByID(ID string) (models.Varian, error)
}

type varianRepository struct {
	db *gorm.DB
}

func NewVarianRepository(db *gorm.DB) *varianRepository {
	return &varianRepository{db}
}

func (r *varianRepository) GetAll() ([]models.Varian, error) {
	var varians []models.Varian

	err := r.db.Table("varian").Select("id", "harga", "jenis").Find(&varians).Error
	return varians, err
}

func (r *varianRepository) GetByID(ID string) (models.Varian, error) {
	var varian models.Varian

	err := r.db.Table("varian").Select("id", "harga", "jenis").Where("id = ?", ID).Find(&varian).Error
	return varian, err
}
