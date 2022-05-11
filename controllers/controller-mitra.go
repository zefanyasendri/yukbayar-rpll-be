package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
)

type mitracontroller struct {
	service services.MitraService
}

func NewMitraController(service services.MitraService) *mitracontroller {
	return &mitracontroller{service}
}

func (con *mitracontroller) InsertMitra(c *gin.Context) {
	req := new(models.Mitra)

	if err := c.Bind(req); err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
	}

	mitra, err := con.service.Create(req)

	// if exists {
	// 	helpers.SendErrorResponse(c, helpers.Response{
	// 		Message: "Failed to create mitra",
	// 		Data:    "Mitra already exist",
	// 	})
	// } else
	if err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to create mitra",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Update mitra data success"
		res.Data = mitra
		helpers.SendSuccessResponse(c, res)
	}
}

func (con *mitracontroller) GetAllMitra(c *gin.Context) {
	mitras, err := con.service.GetAll()

	if len(mitras) == 0 {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Empty mitras",
			Data:    err.Error(),
		})
	}

	var res helpers.Response
	res.Message = "Get mitras success"
	res.Data = mitras
	helpers.SendSuccessResponse(c, res)
}
