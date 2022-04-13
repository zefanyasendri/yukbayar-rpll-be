package pengguna

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(pengguna Pengguna) (Pengguna, error)
	GetAll() ([]Pengguna, error)
	GetByID(ID string) (Pengguna, error)
	UpdateByID(ID string, input interface{}) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(pengguna Pengguna) (Pengguna, error) {
	return pengguna, nil //belum ada logic apa2
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Pengguna, error) {
	var penggunas []Pengguna

	err := r.db.Table("pengguna").Find(&penggunas).Error
	return penggunas, err
}

func (r *repository) GetByID(ID string) (Pengguna, error) {
	var pengguna Pengguna

	err := r.db.Table("pengguna").Where("id = ?", ID).Scan(&pengguna).Error
	return pengguna, err
}

func (r *repository) UpdateByID(ID string, input interface{}) error {
	err := r.db.Table("pengguna").Where("id = ?", ID).Updates(input).Error
	return err
}
