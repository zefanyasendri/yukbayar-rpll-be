package kategoriLayanan

import (
	"yukbayar-rpll-be/models/varian"
)

type KategoriLayanan struct {
	ID       int             `form:"id" json:"id"`
	Kategori string          `form:"email" json:"email"`
	Varian   []varian.Varian `form:"varian" json:"varian"`
}
