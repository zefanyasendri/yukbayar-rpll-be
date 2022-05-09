package walletTopUp

import (
	// "yukbayar-rpll-be/modules/pengguna"

	"yukbayar-rpll-be/modules/pengguna"

	"gorm.io/gorm"
)

type topuprepository interface {
	Create(req *WalletTopUp) error
	GetAll() ([]WalletTopUp, error)
	GetDataCount() int64
	GetByPenggunaID(IdPengguna string) ([]WalletTopUp, error)
	GetSaldoByID(ID string) (pengguna.Pengguna, error)
	UpdateSaldoByID(ID string, amount int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(req *WalletTopUp) error {
	err := r.db.Table("wallettopup").Create(req).Error
	return err
}

func (r *repository) GetDataCount() int64 {
	var num int64
	r.db.Table("wallettopup").Count(&num)
	return num
}

func (r *repository) GetAll() ([]WalletTopUp, error) {
	var topups []WalletTopUp
	err := r.db.Table("wallettopup").Find(&topups).Error
	return topups, err
}

func (r *repository) GetByPenggunaID(IdPengguna string) ([]WalletTopUp, error) {
	var topups []WalletTopUp

	err := r.db.Table("wallettopup").Where("idPengguna = ?", IdPengguna).Scan(&topups).Error

	return topups, err
}

func (r *repository) GetSaldoByID(ID string) (pengguna.Pengguna, error) {
	var pengguna pengguna.Pengguna
	err := r.db.Table("pengguna").Select("saldoYukPay").Where("id = ?", ID).Find(&pengguna).Error

	return pengguna, err
}

func (r *repository) UpdateSaldoByID(ID string, amount int) error {
	err := r.db.Table("pengguna").Where("id = ?", ID).Update("saldoYukPay", amount).Error
	return err
}
