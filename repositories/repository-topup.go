package repositories

import (
	"yukbayar-rpll-be/models"

	"gorm.io/gorm"
)

type TopUpRepository interface {
	Create(req *models.TopUp) error
	GetAll() ([]models.TopUpOutput, error)
	GetDataCount() int64
	GetByPenggunaID(IdPengguna string) ([]models.TopUpOutput, error)
	GetSaldoByID(ID string) (models.SaldoPengguna, error)
	UpdateSaldoByID(IDPengguna string, ID string, amount int) error
	GetByID(ID string) (models.TopUpOutput, error)
	GetTelponById(ID string) (models.Pengguna, error)
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

func (r *topUpRepository) GetAll() ([]models.TopUpOutput, error) {
	var topups []models.TopUpOutput
	err := r.db.Table("wallettopup").Select("wallettopup.id, pengguna.nama, wallettopup.kodeYukPay, wallettopup.metode, wallettopup.nominal, wallettopup.tanggal, wallettopup.id_pengguna, wallettopup.status").Joins("join pengguna on wallettopup.id_pengguna=pengguna.id").Scan(&topups).Error
	return topups, err
}

func (r *topUpRepository) GetByPenggunaID(IdPengguna string) ([]models.TopUpOutput, error) {
	var topups []models.TopUpOutput

	err := r.db.Table("wallettopup").Select("wallettopup.id, pengguna.nama, wallettopup.kodeYukPay, wallettopup.metode, wallettopup.nominal, wallettopup.tanggal, wallettopup.id_pengguna, wallettopup.status").Joins("join pengguna on wallettopup.id_pengguna=pengguna.id").Where("id_pengguna = ?", IdPengguna).Scan(&topups).Error

	return topups, err
}

func (r *topUpRepository) GetByID(ID string) (models.TopUpOutput, error) {
	var topups models.TopUpOutput

	err := r.db.Table("wallettopup").Select("wallettopup.id, pengguna.nama, wallettopup.kodeYukPay, wallettopup.metode, wallettopup.nominal, wallettopup.tanggal, wallettopup.id_pengguna, wallettopup.status").Joins("join pengguna on wallettopup.id_pengguna=pengguna.id").Where("id = ?", ID).Find(&topups).Error

	return topups, err
}

func (r *topUpRepository) GetSaldoByID(ID string) (models.SaldoPengguna, error) {
	var pengguna models.SaldoPengguna
	err := r.db.Table("pengguna").Where("id = ?", ID).Find(&pengguna).Error

	return pengguna, err
}

func (r *topUpRepository) GetTelponById(ID string) (models.Pengguna, error) {
	var pengguna models.Pengguna
	err := r.db.Table("pengguna").Select("noTelpon").Where("id = ?", ID).Find(&pengguna).Error
	return pengguna, err
}

func (r *topUpRepository) UpdateSaldoByID(IDPengguna string, ID string, amount int) error {
	err := r.db.Table("pengguna").Where("id = ?", IDPengguna).Update("saldoYukPay", amount).Error
	err = r.db.Table("wallettopup").Where("id_pengguna", IDPengguna).Where("id", ID).Update("status", "Paid").Error
	return err
}
