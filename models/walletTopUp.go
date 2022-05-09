package models

type WalletTopUp struct {
	ID         int     `form:"id" json:"id"`
	KodeYukPay string  `gorm:"column:KodeYukPay" form:"kodeYukPay" json:"kodeYukPay"`
	Metode     string  `form:"metode" json:"metode"`
	Nominal    float32 `form:"nominal" json:"nominal"`
}

type WalletTopUpResponse struct {
	Message string        `form:"message" json:"message"`
	Data    []WalletTopUp `form:"data" json:"data"`
}
