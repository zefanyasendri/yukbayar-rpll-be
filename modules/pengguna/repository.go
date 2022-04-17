package pengguna

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(req *Pengguna) error
	GetAll() ([]Pengguna, error)
	GetByID(ID string) (Pengguna, error)
	GetByEmail(ID string) (bool, error)
	UpdateByID(ID string, req *UpdateRequest) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(req *Pengguna) error {
	err := r.db.Table("pengguna").Create(req).Error
	return err
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

func (r *repository) GetByEmail(email string) (bool, error) {
	var exists bool

	err := r.db.Table("pengguna").Select("count(*) > 0").Where("email = ?", email).Find(&exists).Error
	return exists, err
}

func (r *repository) UpdateByID(ID string, req *UpdateRequest) error {
	err := r.db.Table("pengguna").Where("id = ?", ID).Updates(req).Error
	return err
}
