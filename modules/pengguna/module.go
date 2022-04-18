package pengguna

type Pengguna struct {
	ID           string `gorm:"primary_key" form:"id" json:"id"`
	Email        string `gorm:"type:varchar(255)" form:"email" json:"email"`
	Nama         string `gorm:"type:varchar(255)" form:"nama" json:"nama"`
	NoTelpon     string `gorm:"type:varchar(255);column:noTelpon;default:-" form:"noTelpon" json:"noTelpon"`
	Password     string `gorm:"type:varchar(255)" form:"password" json:"password"`
	TglLahir     string `gorm:"type:varchar(10);column:tglLahir" form:"tglLahir" json:"tglLahir"`
	Gender       string `gorm:"type:varchar(20)" form:"gender" json:"gender"`
	SaldoYukPay  int    `gorm:"type:int;column:saldoYukPay;default:0" form:"saldoYukPay" json:"saldoYukPay"`
	TipePengguna string `gorm:"type:varchar(255);column:tipepengguna;default:Pengguna" form:"tipePengguna" json:"tipePengguna"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type UpdateRequest struct {
	Email    string `form:"email" json:"email"`
	Nama     string `form:"nama" json:"nama"`
	NoTelpon string `gorm:"column:noTelpon" form:"noTelpon" json:"noTelpon"`
	Password string `form:"password" json:"password"`
}

// type LoginResponse struct {
// 	Message string `form:"message" json:"message"`
// 	Type    string `form:"penggunaType" json:"penggunaType"`
// }
