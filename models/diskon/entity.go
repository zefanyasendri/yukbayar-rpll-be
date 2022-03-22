package diskon

type Diskon struct {
	ID         int    `form:"id" json:"id"`
	KodeYukPay string `form:"kodeYukPay" json:"kodeYukPay"`
	Nama       string `form:"nama" json:"nama"`
}
