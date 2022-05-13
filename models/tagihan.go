package models

type Tagihan struct {
	ID_Tagihan  string `form:"id_transaksi" json:"id_transaksi"`
	ID_Pengguna string `form:"id_pengguna" json:"id_pengguna"`
	ID_Kategori string `gorm:"column:id_kategoriLayanan" form:"id_kategori" json:"id_kategori"`
	No_Tagihan  string `gorm:"column:nomor_tagihan" form:"no_tagihan" json:"no_tagihan"`
	Harga       int    `gorm:"column:harga" form:"harga" json:"harga"`
	Status      string `gorm:"default:'Pending'" form:"status" json:"status"`
}
