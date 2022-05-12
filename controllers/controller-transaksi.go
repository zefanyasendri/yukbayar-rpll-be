package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
)

type transaksiController struct {
	transaksiService services.TransaksiService
}

func NewTransaksiController(transaksiService services.TransaksiService) *transaksiController {
	return &transaksiController{transaksiService}
}

func (con *transaksiController) GetTransaksi(c *gin.Context) {
	transactions, err := con.transaksiService.GetAll()
	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error fetching layanan",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Get transaction data success"
		res.Data = transactions
		helpers.SendSuccessResponse(c, res)
	}
}

func (con *transaksiController) GetTransaksiById(c *gin.Context) {
	id := c.Param("id")
	transactions, err := con.transaksiService.GetTransaksiById(id)

	if len(transactions) == 0 {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Transaksi with id =" + " " + id + " " + "not found",
			Data:    nil,
		})
	} else if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error fetching layanan",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Get transaction data success"
		res.Data = transactions
		helpers.SendSuccessResponse(c, res)
	}
}

func (con *transaksiController) GetTotalHarga(c *gin.Context) {
	totalHarga, err := con.transaksiService.GetTotalHarga()
	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error get total harga",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Get total harga success"
		res.Data = totalHarga
		helpers.SendSuccessResponse(c, res)
	}
}
