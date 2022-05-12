package models

type Mitra struct {
	ID             string `form:"id" json:"id"`
	Alamat         string `form:"alamat" json:"alamat"`
	Email          string `form:"email" json:"email"`
	NoTelpon       string `gorm:"type:varchar(255);column:noTelpon" form:"noTelpon" json:"noTelpon"`
	PemilikBisnis  string `gorm:"type:varchar(255);column:pemilikBisnis" form:"pemilikBisnis" json:"pemilikBisnis"`
	BidangUsaha    string `gorm:"type:varchar(255);column:bidangUsaha" form:"bidangUsaha" json:"bidangUsaha"`
	BadanUsaha     string `gorm:"type:varchar(255);column:badanUsaha" form:"badanUsaha" json:"badanUsaha"`
	NamaPerusahaan string `gorm:"type:varchar(255);column:namaPerusahaan" form:"namaPerusahaan" json:"namaPerusahaan"`
}
