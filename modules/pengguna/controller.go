package pengguna

import (
	//"errors"

	//"strconv"

	"yukbayar-rpll-be/helpers/response"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

func (con *controller) GetUserData(c *gin.Context) {
	id := c.Param("id")

	pengguna, err := con.service.GetByID(id)
	if err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "User ID not found",
			Data:    err.Error(),
		})
		return
	}

	var res response.Response
	res.Message = "Get user data success"
	res.Data = pengguna
	response.SendSuccessResponse(c, res)
}

func (con *controller) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := new(UpdateRequest)
	if err := c.Bind(req); err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "Failed to update",
			Data:    err.Error(),
		})
		return
	}

	err := con.service.UpdateByID(id, req)

	if err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "Failed to update",
			Data:    err.Error(),
		})
		return
	}

	pengguna, err := con.service.GetByID(id)
	var res response.Response

	res.Message = "Update user data success"
	res.Data = pengguna
	response.SendSuccessResponse(c, res)
}
