package models

type Layanan struct {
	ID       string `form:"id" json:"id"`
	ID_Mitra string `gorm:"type:varchar(6);column:id_mitra" form:"IdMitra" json:"IdMitra"`
	Nama     string `form:"nama" json:"nama"`
}

/* Contoh
1, PLN
2, Internet
*/
