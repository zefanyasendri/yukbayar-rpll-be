package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
)

type kategoriLayananController struct {
	kategoriLayananService services.KategoriLayananService
}

func NewKategoriLayananController(kategoriLayananService services.KategoriLayananService) *kategoriLayananController {
	return &kategoriLayananController{kategoriLayananService}
}

func (con *kategoriLayananController) GetKategori(c *gin.Context) {
	id := c.Param("id")

	kategoriLayanan, err := con.kategoriLayananService.GetByID(id)
	kategoriLayanan.Varian, err = con.kategoriLayananService.GetIDByKategori(id)

	if kategoriLayanan.ID == "" {
		helpers.SendNoContentResponse(c, helpers.Response{
			Message: "Kategori ID not found",
			Data:    nil,
		})
	} else if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error fetching kategori",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Get kategori data success"
		res.Data = kategoriLayanan
		helpers.SendSuccessResponse(c, res)
	}
}
