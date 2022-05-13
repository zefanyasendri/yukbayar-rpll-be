package models

type Diskon struct {
	ID             int     `form:"id" json:"id"`
	Nama           string  `form:"nama" json:"nama"`
	KodeYukPay     string  `gorm:"column:kode" form:"kodeYukPay" json:"kodeYukPay"`
	JumlahPotongan float64 `gorm:"column:jumlahPotongan" form:"jumlahPotongan" json:"jumlahPotongan"`
}
