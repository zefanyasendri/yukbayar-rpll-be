package models

type TopUp struct {
	ID         string   `gorm:"primary_key" form:"id" json:"id"`
	KodeYukPay string   `gorm:"type:varchar(255);column:kodeYukPay" form:"kodeYukPay" json:"kodeYukPay"`
	Metode     string   `gorm:"type:varchar(255)" form:"metode" json:"metode"`
	Nominal    float32  `gorm:"type:int" form:"nominal" json:"nominal"`
	IdPengguna string   `gorm:"type:varchar(36);column:id_pengguna" form:"IdPengguna" json:"IdPengguna"`
	Pengguna   Pengguna `gorm:"foreignKey:IdPengguna;references:ID"`
}

// type WalletTopUpResponse struct {
// 	Message string        `form:"message" json:"message"`
// 	Data    []WalletTopUp `form:"data" json:"data"`
// }
