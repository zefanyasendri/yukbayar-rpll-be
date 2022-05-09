package models

type KategoriLayanan struct {
	ID       string   `form:"id" json:"id"`
	Kategori string   `form:"kategori" json:"kategori"`
	Varian   []Varian `gorm:"references:Varian.ID" form:"varian" json:"varian"`
}

/* Contoh
11, PLN-Token, 1
12, PLN-Tagihan, 1
21, Inet-Indihome, 2
22, Inet-MyRep, 2
23, Inet-Biznet, 2
*/
