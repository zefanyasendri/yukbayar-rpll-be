package models

type Varian struct {
	ID    string `form:"id" json:"id"`
	Harga int    `form:"harga" json:"harga"`
	Jenis string `form:"jenis" json:"jenis"`
}

/*
Contoh:
111, 20000, PLN-Token20k, 11
112, 50000, PLN-Token50k, 11
113, 100000, PLN-Token100k, 11
121, 90000, PLN-Tagihan, 12
211, 100000, Indihome-Tagiha, 21
221, 100000, MyRep-Tagihan, 22
231, 100000, BizNet-Tagihan, 23
*/
