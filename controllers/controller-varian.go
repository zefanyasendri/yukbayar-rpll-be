package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
)

type varianController struct {
	varianService services.VarianService
}

func NewVarianController(varianService services.VarianService) *varianController {
	return &varianController{varianService}
}

func (con *varianController) GetVarian(c *gin.Context) {
	id := c.Param("id")

	varian, err := con.varianService.GetByID(id)
	if varian.ID == "" {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Varian ID not found",
			Data:    nil,
		})
	} else if err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Error fetching varian",
			Data:    nil,
		})
	} else {
		var res helpers.Response
		res.Message = "Get varian data success"
		res.Data = varian
		helpers.SendSuccessResponse(c, res)
	}

}
