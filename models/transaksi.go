package models

type Transaksi struct {
	ID         int        `form:"id" json:"id"`
	Pengguna   []Pengguna `form:"pengguna" json:"pengguna"`
	TotalHarga int        `form:"totalHarga" json:"totalHarga"`
	Layanan    []Layanan  `form:"layanan" json:"layanan"`
}

type TransaksiResponse struct {
	Message string      `form:"message" json:"message"`
	Data    []Transaksi `form:"data" json:"data"`
}
