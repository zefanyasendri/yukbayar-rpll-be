package controllers

import (
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/models"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
)

type penggunaController struct {
	penggunaService services.PenggunaService
}

func NewPenggunaController(penggunaService services.PenggunaService) *penggunaController {
	return &penggunaController{penggunaService}
}

func (con *penggunaController) GetUserData(c *gin.Context) {
	id := c.Param("id")

	pengguna, err := con.penggunaService.GetByID(id)
	if err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "User ID not found",
			Data:    err.Error(),
		})
		return
	}

	var res helpers.Response
	res.Message = "Get user data success"
	res.Data = pengguna
	helpers.SendSuccessResponse(c, res)
}

func (con *penggunaController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := new(models.PenggunaUpdateRequest)
	if err := c.Bind(req); err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
		return
	}

	err := con.penggunaService.UpdateByID(id, req)

	if err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to update",
			Data:    err.Error(),
		})
		return
	}

	pengguna, err := con.penggunaService.GetByID(id)
	var res helpers.Response

	res.Message = "Update user data success"
	res.Data = pengguna
	helpers.SendSuccessResponse(c, res)
}

func (con *penggunaController) Register(c *gin.Context) {
	req := new(models.Pengguna)

	if err := c.Bind(req); err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
	}

	uuid, isExist, err := con.penggunaService.Create(req)

	if isExist {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to create account",
			Data:    "Email already exist",
		})
	} else if err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to create account",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Update user data success"
		res.Data = map[string]interface{}{"uuid": uuid}
		helpers.SendSuccessResponse(c, res)
	}
}

func (con *penggunaController) Login(c *gin.Context) {
	req := new(models.PenggunaLoginRequest)

	if err := c.Bind(req); err != nil {
		helpers.SendErrorResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
	}

	//token, err := "token"

	var res helpers.Response
	res.Message = "Login user success"
	res.Data = "some fucking jwt token or some shit."
	helpers.SendSuccessResponse(c, res)
}
