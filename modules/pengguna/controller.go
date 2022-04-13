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
	email := c.PostForm("email")
	nama := c.PostForm("nama")
	noTelpon := c.PostForm("noTelpon")
	password := c.PostForm("password")
	// tglLahir := c.PostForm("tglLahir")
	// gender = c.PostForm("gender")

	id := c.Param("id")

	pengguna, err := con.service.GetByID(id)
	if err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "User ID Not Found",
			Data:    err.Error(),
		})
		return
	}

	//logic ini harusnya di service tapi saya harus mikir dlu
	if email == "" {
		email = pengguna.Email
	}

	if nama == "" {
		nama = pengguna.Nama
	}

	if noTelpon == "" {
		noTelpon = pengguna.NoTelpon
	}

	if password == "" {
		password = pengguna.Password
	}

	result := con.service.UpdateByID(id, map[string]interface{}{"email": email, "nama": nama, "noTelpon": noTelpon, "password": password})

	if err != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "Failed to update",
			Data:    err.Error(),
		})
		return
	}

	var res response.Response

	res.Message = "Update user data success"
	res.Data = result
	response.SendSuccessResponse(c, res)
}

// func CheckOldPass(c *gin.Context) {

// 	db := repo.Connect()
// 	defer db.Close()

// 	query := "SELECT password from pengguna"

// 	id := c.Param("id")
// 	pass := c.PostForm("password")

// 	if id != "" {
// 		query += "WHERE id='" + id + "'"
// 	}

// 	var pengguna model.Pengguna
// 	if pass != pengguna.Password {
// 		err := errors.New("different pass")
// 		log.Print("Error: ", err)
// 	}

// }
