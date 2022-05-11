package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type PenggunaRepository interface {
	Create(req *models.Pengguna) error
	GetAll() ([]models.PenggunaResponse, error)
	GetByID(ID string) (models.Pengguna, error)
	GetByEmail(ID string) (models.Pengguna, error)
	UpdateByID(ID string, req *models.PenggunaUpdateRequest) error
}

type penggunaRepository struct {
	db *gorm.DB
}

func NewPenggunaRepository(db *gorm.DB) *penggunaRepository {
	return &penggunaRepository{db}
}

func (r *penggunaRepository) Create(req *models.Pengguna) error {
	err := r.db.Table("pengguna").Create(req).Error
	return err
}

func (r *penggunaRepository) GetAll() ([]models.PenggunaResponse, error) {
	var penggunas []models.PenggunaResponse

	err := r.db.Table("pengguna").Select("id", "nama", "email", "noTelpon", "tglLahir", "gender").Find(&penggunas).Error
	return penggunas, err
}

func (r *penggunaRepository) GetByID(ID string) (models.Pengguna, error) {
	var pengguna models.Pengguna

	err := r.db.Table("pengguna").Where("id = ?", ID).Scan(&pengguna).Error
	return pengguna, err
}

func (r *penggunaRepository) GetByEmail(email string) (models.Pengguna, error) {
	var pengguna models.Pengguna

	err := r.db.Table("pengguna").Where("email = ?", email).Scan(&pengguna).Error
	return pengguna, err
}

func (r *penggunaRepository) UpdateByID(ID string, req *models.PenggunaUpdateRequest) error {
	err := r.db.Table("pengguna").Where("id = ?", ID).Updates(req).Error
	return err
}
