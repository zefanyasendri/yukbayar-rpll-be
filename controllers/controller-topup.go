package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/services"

	// "yukbayar-rpll-be/modules/pengguna"

	"github.com/gin-gonic/gin"
)

type topUpController struct {
	topUpService services.TopUpService
}

func NewTopUpController(topUpService services.TopUpService) *topUpController {
	return &topUpController{topUpService}
}

func (con *topUpController) InsertNewTopup(c *gin.Context) {

	req := new(models.TopUp)

	if err := c.Bind(req); err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
	}

	id, kode, err := con.topUpService.Create(req)

	if err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to insert data",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Insert TopUp Data success"
		res.Data = map[string]interface{}{"id": id, "kodeYukPay": kode}
		helpers.SendSuccessResponse(c, res)
	}

}

func (con *topUpController) UpdateSaldoUser(c *gin.Context) {
	id := c.Param("id")
	req := new(models.Pengguna)

	if err := c.Bind(req); err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
		return
	}

	updatedAmount, err := con.topUpService.UpdateSaldoByID(id, req.SaldoYukPay)

	if err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to update",
			Data:    err.Error(),
		})
		return
	}
	var res helpers.Response

	res.Message = "Update user data success"
	res.Data = updatedAmount
	helpers.SendSuccessResponse(c, res)

}
