package pengguna

type Pengguna struct {
	ID           int    `form:"id" json:"id"`
	Email        string `form:"email" json:"email"`
	Nama         string `form:"nama" json:"nama"`
	NoTelpon     string `form:"noTelpon" json:"noTelpon"`
	Password     string `form:"password" json:"password"`
	TglLahir     string `form:"tglLahir" json:"tglLahir"`
	Gender       string `form:"gender" json:"gender"`
	SaldoYukPay  int    `form:"saldoYukPay" json:"saldoYukPay"`
	TipePengguna string `form:"tipepengguna" json:"tipepengguna"`
}

// type PenggunaResponse struct {
// 	Message string     `form:"message" json:"message"`
// 	Data    []Pengguna `form:"data" json:"data"`
// }

// type LoginResponse struct {
// 	Message string `form:"message" json:"message"`
// 	Type    string `form:"penggunaType" json:"penggunaType"`
// }
