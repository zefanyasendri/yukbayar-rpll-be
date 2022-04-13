package transaksi

import (
	"yukbayar-rpll-be/modules/layanan"
	"yukbayar-rpll-be/modules/pengguna"
)

type Transaksi struct {
	ID         int                 `form:"id" json:"id"`
	Pengguna   []pengguna.Pengguna `form:"pengguna" json:"pengguna"`
	TotalHarga int                 `form:"totalHarga" json:"totalHarga"`
	Layanan    []layanan.Layanan   `form:"layanan" json:"layanan"`
}

type TransaksiResponse struct {
	Message string      `form:"message" json:"message"`
	Data    []Transaksi `form:"data" json:"data"`
}