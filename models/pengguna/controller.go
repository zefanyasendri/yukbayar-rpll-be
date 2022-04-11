package pengguna

import (
	//"errors"

	//"strconv"

	"fmt"
	"yukbayar-rpll-be/db"
	"yukbayar-rpll-be/helpers/response"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	db := db.Connect()
	id := c.Param("id")

	var pengguna Pengguna

	result := db.First(&pengguna, id)
	if result.Error != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "User ID not found",
			Data:    result.Error,
		})
		return
	}

	var res response.Response
	fmt.Println(pengguna.NoTelpon)
	res.Message = "Get user data success"
	res.Data = result.RowsAffected
	response.SendSuccessResponse(c, res)
}

func UpdateUser(c *gin.Context) {
	db := db.Connect()

	email := c.PostForm("email")
	nama := c.PostForm("nama")
	noTelpon := c.PostForm("noTelpon")
	password := c.PostForm("password")
	// tglLahir := c.PostForm("tglLahir")
	// gender = c.PostForm("gender")

	id := c.Param("id")

	var pengguna Pengguna

	queryRes := db.Table("pengguna").Select("email, nama, noTelpon, password").Where("id = ?", id).Find(&pengguna)
	if queryRes.Error != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "User ID not found",
		})
		return
	}

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

	updateRes := db.Table("pengguna").Where("id = ?", id).Updates(map[string]interface{}{"email": email, "nama": nama, "noTelpon": noTelpon, "password": password})

	if updateRes.Error != nil {
		response.SendErrorResponse(c, response.Response{
			Message: "Failed to update",
		})
		return
	}

	var res response.Response

	res.Message = "Update user data success"
	res.Data = updateRes.RowsAffected
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
