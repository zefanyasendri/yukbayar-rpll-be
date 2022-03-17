package model

type Pengguna struct {
	ID          int    `form:"id" json:"id"`
	Email       string `form:"email" json:"email"`
	Nama        string `form:"nama" json:"nama"`
	NoTelpon    string `form:"noTelpon" json:"noTelpon"`
	Password    string `form:"password" json:"password"`
	TglLahir    string `form:"password" json:"tglLahir"`
	Gender      string `form:"password" json:"gender"`
	SaldoYukPay int    `form:"saldoYukPay" json:"saldoYukPay"`
	// tipePengguna: TipePengguna

}

type PenggunaResponse struct {
	Message string     `form:"message" json:"message"`
	Data    []Pengguna `form:"data" json:"data"`
}

type LoginResponse struct {
	Message string `form:"message" json:"message"`
	Type    string `form:"penggunaType" json:"penggunaType"`
}
