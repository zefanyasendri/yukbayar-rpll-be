package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type TagihanRepository interface {
	Create(req *models.Tagihan) (models.Tagihan, error)
	GetPendingStatus(IdKategori string, IdPengguna string) ([]models.Tagihan, error)
	UpdateStatus(IdTagihan string) (models.Tagihan, error)
	GetDataCount() int64
}

type tagihanRepository struct {
	db *gorm.DB
}

func NewTagihanRepository(db *gorm.DB) *tagihanRepository {
	return &tagihanRepository{db}
}

func (r *tagihanRepository) Create(req *models.Tagihan) (models.Tagihan, error) {
	err := r.db.Table("tagihan").Create(req).Error

	var tagihan models.Tagihan
	err = r.db.Table("tagihan").Where("id_tagihan = ?", req.ID_Tagihan).Error
	return tagihan, err
}

func (r *tagihanRepository) GetDataCount() int64 {
	var num int64
	r.db.Table("tagihan").Count(&num)
	return num
}

func (r *tagihanRepository) GetPendingStatus(IdKategori string, IdPengguna string) ([]models.Tagihan, error) {
	var tagihans []models.Tagihan
	err := r.db.Table("tagihan").Select("tagihan.id_tagihan, pengguna.id, kategorilayanan.id, tagihan.nomor_tagihan, tagihan.harga, tagihan.status").Joins("join kategorilayanan on tagihan.id_kategoriLayanan=kategorilayanan.id").Joins("join pengguna on tagihan.id_pengguna=pengguna.id").Where("status = ?", "Pending").Scan(&tagihans).Error
	return tagihans, err
}

func (r *tagihanRepository) UpdateStatus(IdTagihan string) (models.Tagihan, error) {
	err := r.db.Table("tagihan").Where("id_tagihan = ?", IdTagihan).Update("status", "Paid").Error

	var tagihan models.Tagihan
	err = r.db.Table("tagihan").Select("tagihan.id_tagihan, pengguna.id, kategorilayanan.id, tagihan.nomor_tagihan, tagihan.harga, tagihan.status").Joins("join kategorilayanan on tagihan.id_kategoriLayanan=kategorilayanan.id").Joins("join pengguna on tagihan.id_pengguna=pengguna.id").Find(&tagihan).Error
	return tagihan, err
}
