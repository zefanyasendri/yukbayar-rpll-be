package models

import "time"

type TopUp struct {
	ID string `gorm:"primary_key" form:"id" json:"id"`
	//IdPengguna   string    `gorm:"type:varchar(36);column:id_pengguna" form:"IdPengguna" json:"IdPengguna"`
	ID_Pengguna string `form:"id_pengguna" json:"id_pengguna"`
	//Pengguna     string    `gorm:"column:nama" form:"nama" json:"nama"`
	KodeYukPay   string    `gorm:"type:varchar(255);column:kodeYukPay" form:"kodeYukPay" json:"kodeYukPay"`
	Metode       string    `gorm:"type:varchar(255)" form:"metode" json:"metode"`
	Nominal      float32   `gorm:"type:int" form:"nominal" json:"nominal"`
	TanggalTopUp time.Time `gorm:"column:tanggal" form:"tanggal" json:"tanggal"`
	Status       string    `form:"status" json:"status"`
	//Pengguna     Pengguna  `gorm:"foreignKey:ID_Pengguna;references:ID"`
}

/* Contoh Input
eBanking-BCA, 100000, *Insert ID_Pengguna, Belum Bayar
*/
