package layanan

import (
	"yukbayar-rpll-be/modules/kategoriLayanan"
)

type Layanan struct {
	ID              int                               `form:"id" json:"id"`
	Jenis           string                            `form:"jenis" json:"jenis"`
	KategoriLayanan []kategoriLayanan.KategoriLayanan `form:"kategoriLayanan" json:"kategoriLayanan"`
	Nama            string                            `form:"nama" json:"nama"`
}
