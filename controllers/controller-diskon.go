package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
)

type diskonController struct {
	diskonService services.DiskonService
}

func NewDiskonController(diskonService services.DiskonService) *diskonController {
	return &diskonController{diskonService}
}

func (con *diskonController) GetAll(c *gin.Context) {
	diskons, err := con.diskonService.GetAll()

	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error fetching diskon",
			Data:    err.Error(),
		})
		return
	}

	var res helpers.Response
	res.Message = "Get all diskon success"
	res.Data = diskons
	helpers.SendSuccessResponse(c, res)
}
