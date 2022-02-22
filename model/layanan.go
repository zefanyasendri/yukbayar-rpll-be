package model

type Layanan struct {
	ID              int               `form:"id" json:"id"`
	Jenis           string            `form:"jenis" json:"jenis"`
	KategoriLayanan []KategoriLayanan `form:"kategoriLayanan" json:"kategoriLayanan"`
	Nama            string            `form:"nama" json:"nama"`
}

type LayananResponse struct {
	Message string    `form:"message" json:"message"`
	Data    []Layanan `form:"data" json:"data"`
}
