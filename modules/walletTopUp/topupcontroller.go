package walletTopUp

import (
	"yukbayar-rpll-be/helpers/response"
	"yukbayar-rpll-be/modules/pengguna"
	// "yukbayar-rpll-be/modules/pengguna"

	"github.com/gin-gonic/gin"
)

type topupcontroller struct {
	service topupservice
}

func NewController(service topupservice) *topupcontroller {
	return &topupcontroller{service}
}

func (con *topupcontroller) InsertNewTopup(c *gin.Context) {

	req := new(WalletTopUp)

	if err := c.Bind(req); err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
	}

	id, kode, err := con.service.Create(req)

	if err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "Failed to insert data",
			Data:    err.Error(),
		})
	} else {
		var res response.Response
		res.Message = "Insert TopUp Data success"
		res.Data = map[string]interface{}{"id": id, "kodeYukPay": kode}
		response.SendSuccessResponse(c, res)
	}

}

func (con *topupcontroller) UpdateSaldoUser(c *gin.Context) {
	id := c.Param("id")
	req := new(pengguna.Pengguna)

	if err := c.Bind(req); err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
		return
	}

	updatedAmount, err := con.service.UpdateSaldoByID(id, req.SaldoYukPay)

	if err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "Failed to update",
			Data:    err.Error(),
		})
		return
	}
	var res response.Response

	res.Message = "Update user data success"
	res.Data = updatedAmount
	response.SendSuccessResponse(c, res)

}
