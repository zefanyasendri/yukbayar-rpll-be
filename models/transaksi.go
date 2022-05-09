package models

import "time"

type Transaksi struct {
	ID_Transaksi     string    `form:"id_transaksi" json:"id_transaksi"`
	ID_Pengguna      string    `form:"id_pengguna" json:"id_pengguna"`
	Pengguna         string    `gorm:"column:nama" form:"nama" json:"nama"`
	Varian           string    `gorm:"column:jenis" form:"varian" json:"varian"`
	TotalHarga       int       `gorm:"column:harga" form:"totalHarga" json:"totalHarga"`
	TanggalTransaksi time.Time `gorm:"column:tanggal" form:"tanggal" json:"tanggal"`
	Status           string    `form:"status" json:"status"`
}

/*
contoh
INET-1, uuid-kont-1, Benzefa, Internet
*/
