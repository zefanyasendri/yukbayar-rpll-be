package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
)

type tagihanController struct {
	tagihanService services.TagihanService
}

func NewTagihanController(tagihanService services.TagihanService) *tagihanController {
	return &tagihanController{tagihanService}
}
func (con *tagihanController) Create(c *gin.Context) {
	req := new(models.Tagihan)

	if err := c.Bind(req); err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
	}

	tagihan, err := con.tagihanService.Create(req)

	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Failed to insert data",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Insert Tagihan Data success"
		res.Data = tagihan
		helpers.SendSuccessResponse(c, res)
	}
}

func (con *tagihanController) GetPendingStatus(c *gin.Context) {
	idKategori := c.Param("idKategori")
	idPengguna := c.Param("idPengguna")
	tagihans, err := con.tagihanService.GetPendingStatus(idKategori, idPengguna)

	if len(tagihans) == 0 {
		helpers.SendNotFoundResponse(c, helpers.Response{
			Message: "Empty Data",
			Data:    nil,
		})
	} else if err != nil {
		helpers.SendNotFoundResponse(c, helpers.Response{
			Message: "Cannot retrive tagihan data",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Get topup data success"
		res.Data = tagihans
		helpers.SendSuccessResponse(c, res)
	}
}

func (con *tagihanController) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	tagihan, err := con.tagihanService.UpdateStatus(id)

	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error Update Status",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Update Success"
		res.Data = tagihan
		helpers.SendSuccessResponse(c, res)
	}
}
