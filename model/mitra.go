package model

type Mitra struct {
	ID            int       `form:"id" json:"id"`
	Alamat        string    `form:"alamat" json:"alamat"`
	Email         string    `form:"email" json:"email"`
	Layanan       []Layanan `form:"layanan" json:"layanan"`
	NoTelpon      string    `form:"noTelpon" json:"noTelpon"`
	PemilikBisnis string    `form:"pemilikBisnis" json:"pemilikBisnis"`
}

type MitraResponse struct {
	Message string  `form:"message" json:"message"`
	Data    []Mitra `form:"data" json:"data"`
}
