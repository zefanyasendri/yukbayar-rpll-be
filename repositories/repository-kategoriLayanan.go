package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type KategoriLayananRepository interface {
	GetAll() ([]models.KategoriLayanan, error)
	GetByID(ID string) (models.KategoriLayanan, error)
	GetIDByKategori(ID string) ([]models.Varian, error)
}

type kategoriLayananRepository struct {
	db *gorm.DB
}

func NewKategoriLayananRepository(db *gorm.DB) *kategoriLayananRepository {
	return &kategoriLayananRepository{db}
}

func (r *kategoriLayananRepository) GetAll() ([]models.KategoriLayanan, error) {
	var kategoriLayanans []models.KategoriLayanan

	err := r.db.Table("varian").Find(&kategoriLayanans).Error
	return kategoriLayanans, err
}

func (r *kategoriLayananRepository) GetByID(ID string) (models.KategoriLayanan, error) {
	var kategoriLayanan models.KategoriLayanan

	err := r.db.Table("kategorilayanan").Select("id", "kategori").Where("id = ?", ID).Scan(&kategoriLayanan).Error
	return kategoriLayanan, err
}

func (r *kategoriLayananRepository) GetIDByKategori(ID string) ([]models.Varian, error) {
	var varians []models.Varian

	err := r.db.Table("varian").Select("id", "harga", "jenis").Where("id_kategori = ?", ID).Scan(&varians).Error
	return varians, err
}
