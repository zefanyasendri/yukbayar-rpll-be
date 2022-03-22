package varian

type Varian struct {
	ID    int    `form:"id" json:"id"`
	Harga int    `form:"harga" json:"harga"`
	Jenis string `form:"jenis" json:"jenis"`
}

type VarianResponse struct {
	Message string   `form:"message" json:"message"`
	Data    []Varian `form:"data" json:"data"`
}
