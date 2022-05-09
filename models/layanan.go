package models

type Layanan struct {
	ID   string `form:"id" json:"id"`
	Nama string `form:"nama" json:"nama"`
}

/* Contoh
1, PLN
2, Internet
*/
