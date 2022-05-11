package repositories

import (
	// "yukbayar-rpll-be/modules/pengguna"

	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type TopUpRepository interface {
	Create(req *models.TopUp) error
	GetAll() ([]models.TopUp, error)
	GetDataCount() int64
	GetByPenggunaID(IdPengguna string) ([]models.TopUp, error)
	GetSaldoByID(ID string) (models.Pengguna, error)
	UpdateSaldoByID(ID string, amount int) error
	GetByID(ID string) ([]models.TopUp, error)
}

type topUpRepository struct {
	db *gorm.DB
}

func NewTopUpRepository(db *gorm.DB) *topUpRepository {
	return &topUpRepository{db}
}

func (r *topUpRepository) Create(req *models.TopUp) error {
	err := r.db.Table("wallettopup").Create(req).Error
	return err
}

func (r *topUpRepository) GetDataCount() int64 {
	var num int64
	r.db.Table("wallettopup").Count(&num)
	return num
}

func (r *topUpRepository) GetAll() ([]models.TopUp, error) {
	var topups []models.TopUp
	err := r.db.Table("wallettopup").Find(&topups).Error
	return topups, err
}

func (r *topUpRepository) GetByPenggunaID(IdPengguna string) ([]models.TopUp, error) {
	var topups []models.TopUp

	err := r.db.Table("wallettopup").Where("idPengguna = ?", IdPengguna).Scan(&topups).Error

	return topups, err
}

func (r *topUpRepository) GetByID(ID string) ([]models.TopUp, error) {
	var topups []models.TopUp

	err := r.db.Table("wallettopup").Where("id = ?", ID).Scan(&topups).Error

	return topups, err
}

func (r *topUpRepository) GetSaldoByID(ID string) (models.Pengguna, error) {
	var pengguna models.Pengguna
	err := r.db.Table("pengguna").Select("saldoYukPay").Where("id = ?", ID).Find(&pengguna).Error

	return pengguna, err
}

func (r *topUpRepository) UpdateSaldoByID(ID string, amount int) error {
	err := r.db.Table("pengguna").Where("id = ?", ID).Update("saldoYukPay", amount).Error
	return err
}
