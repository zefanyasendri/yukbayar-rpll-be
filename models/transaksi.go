package models

import "time"

type Transaksi struct {
	ID_Transaksi     string    `form:"id_transaksi" json:"id_transaksi"`
	ID_Pengguna      string    `form:"id_pengguna" json:"id_pengguna"`
	Pengguna         string    `gorm:"column:nama" form:"nama" json:"nama"`
	Varian           string    `gorm:"column:jenis" form:"varian" json:"varian"`
	TotalHarga       int       `gorm:"column:totalHarga" form:"totalHarga" json:"totalHarga"`
	TanggalTransaksi time.Time `gorm:"column:tanggal" form:"tanggal" json:"tanggal"`
	Status           string    `form:"status" json:"status"`
}

type NewTransaksi struct {
	ID_Transaksi string `json:"id_transaksi"`
	ID_Pengguna  string `form:"id_pengguna" json:"id_pengguna"`
	TotalHarga   int    `gorm:"column:totalHarga" form:"totalHarga" json:"totalHarga"`
	ID_varian    string `form:"id_varian" json:"id_varian"`
}

/*
contoh
1, uuid-kont-1, Benzefa, PDAM memek, 102000, 12390123-12390123-123-91091,
*/
