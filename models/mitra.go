package models

type Mitra struct {
	ID string `form:"id" json:"id"`
	//ID_Layanan    string    `gorm:"type:varchar(6);column:id_layanan" form:"IdLayanan" json:"IdLayanan"`
	Alamat        string `form:"alamat" json:"alamat"`
	Email         string `form:"email" json:"email"`
	NoTelpon      string `gorm:"type:varchar(255);column:noTelpon" form:"noTelpon" json:"noTelpon"`
	PemilikBisnis string `gorm:"type:varchar(255);column:pemilikBisnis" form:"pemilikBisnis" json:"pemilikBisnis"`
	BidangUsaha   string `gorm:"type:varchar(255);column:bidangUsaha" form:"bidangUsaha" json:"bidangUsaha"`
	//Layanan       []Layanan `form:"layanan" json:"layanan"`
}
