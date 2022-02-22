package model

type Diskon struct {
	ID         int    `form:"id" json:"id"`
	KodeYukPay string `form:"kodeYukPay" json:"kodeYukPay"`
	Nama       string `form:"nama" json:"nama"`
}

type DiskonResponse struct {
	Message string   `form:"message" json:"message"`
	Data    []Diskon `form:"data" json:"data"`
}
