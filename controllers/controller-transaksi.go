package controllers

import (
	"fmt"
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
	fmt.Println(transactions[0].Pengguna)
	if err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
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
