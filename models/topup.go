package models

import "time"

type TopUp struct {
	ID           string    `gorm:"primary_key" form:"id" json:"id"`
	KodeYukPay   string    `gorm:"type:varchar(255);column:kodeYukPay" form:"kodeYukPay" json:"kodeYukPay"`
	Metode       string    `gorm:"type:varchar(255)" form:"metode" json:"metode"`
	Nominal      float32   `gorm:"type:int" form:"nominal" json:"nominal"`
	IdPengguna   string    `gorm:"type:varchar(36);column:id_pengguna" form:"IdPengguna" json:"IdPengguna"`
	TanggalTopUp time.Time `gorm:"column:tanggal" form:"tanggal" json:"tanggal"`
	Pengguna     Pengguna  `gorm:"foreignKey:IdPengguna;references:ID"`
}

// type WalletTopUpResponse struct {
// 	Message string        `form:"message" json:"message"`
// 	Data    []WalletTopUp `form:"data" json:"data"`
// }

/* Contoh Input
eBanking-BCA, 100000, f8bc4854-71c2-4bf6-85a8-95bbf38f
*/
