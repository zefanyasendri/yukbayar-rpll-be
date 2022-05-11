package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
)

type layananController struct {
	layananService services.LayananService
}

func NewLayananController(layananService services.LayananService) *layananController {
	return &layananController{layananService}
}

func (con *layananController) GetLayanan(c *gin.Context) {

	layanans, err := con.layananService.GetAll()

	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error fetching layanan",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Get kategori data success"
		res.Data = layanans
		helpers.SendSuccessResponse(c, res)
	}
}
