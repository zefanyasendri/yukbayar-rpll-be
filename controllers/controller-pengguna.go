package controllers

import (
	"fmt"
	"strconv"
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
		helpers.SendNotFoundResponse(c, helpers.Response{
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

func (con *penggunaController) GetAllData(c *gin.Context) {
	penggunas, err := con.penggunaService.GetAll()

	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Cannot retrive user data",
			Data:    err.Error(),
		})
		return
	}

	var res helpers.Response
	res.Message = "Get user data success"
	res.Data = penggunas
	helpers.SendSuccessResponse(c, res)
}

func (con *penggunaController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := new(models.PenggunaUpdateRequest)
	if err := c.Bind(req); err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
		return
	}
	fmt.Println(req)

	pengguna, match, err := con.penggunaService.UpdateByID(id, req)
	var res helpers.Response

	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Failed to update",
			Data:    err.Error(),
		})
		return
	}
	if !match {
		if req.OldPassword == "" && req.Password == "" {
			res.Message = "Update user data success"
			res.Data = pengguna
			helpers.SendSuccessResponse(c, res)
		} else {
			helpers.SendBadRequestResponse(c, helpers.Response{
				Message: "Password incorrect",
				Data:    nil,
			})
		}
		return
	} else {
		if req.Password == "" {
			helpers.SendBadRequestResponse(c, helpers.Response{
				Message: "You didn't put a new password",
				Data:    nil,
			})
		} else {
			res.Message = "Update user data success"
			res.Data = pengguna
			helpers.SendSuccessResponse(c, res)
		}
		return
	}

}

func (con *penggunaController) Register(c *gin.Context) {
	req := new(models.Pengguna)

	if err := c.Bind(req); err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
	}

	uuid, isExist, err := con.penggunaService.Create(req)

	if isExist {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Failed to create account",
			Data:    "Email already exist",
		})
	} else if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
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
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Failed to bind",
			Data:    err.Error(),
		})
	}

	pengguna, match, err := con.penggunaService.GetAccount(req.Email, req.Password)

	if pengguna.Email == "" {
		helpers.SendNotFoundResponse(c, helpers.Response{
			Message: "User ID not found",
			Data:    nil,
		})
	} else if !match {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Password doesn't match",
			Data:    nil,
		})
	} else if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error retrieving pengguna",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Welcome!"
		res.Data = pengguna
		helpers.SendSuccessResponse(c, res)
	}
}

func (con *penggunaController) UpdateSaldo(c *gin.Context) {
	id := c.Param("id")
	saldo, _ := strconv.Atoi(c.PostForm("saldo"))

	pengguna, err := con.penggunaService.UpdateSaldoByID(id, saldo)
	if err != nil {
		helpers.SendBadRequestResponse(c, helpers.Response{
			Message: "Error updating saldo",
			Data:    err.Error(),
		})
	} else {
		var res helpers.Response
		res.Message = "Update saldo successful"
		res.Data = pengguna
		helpers.SendSuccessResponse(c, res)
	}
}
