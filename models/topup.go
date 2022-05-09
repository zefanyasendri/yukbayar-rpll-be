package models

type TopUp struct {
	ID         int    `form:"id" json:"id"`
	KodeYukPay string `form:"kodeYukPay" json:"kodeYukPay"`
	Metode     string `form:"metode" json:"metode"`
	Nominal    int    `form:"nominal" json:"nominal"`
}

type TopUpResponse struct {
	Message string  `form:"message" json:"message"`
	Data    []TopUp `form:"data" json:"data"`
}
