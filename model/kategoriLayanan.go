package model

type KategoriLayanan struct {
	ID       int      `form:"id" json:"id"`
	Kategori string   `form:"email" json:"email"`
	Varian   []Varian `form:"varian" json:"varian"`
}

type KategoriLayananResponse struct {
	Message string            `form:"message" json:"message"`
	Data    []KategoriLayanan `form:"data" json:"data"`
}
