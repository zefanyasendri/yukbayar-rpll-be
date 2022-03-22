package pengguna

import (
	//"errors"

	"database/sql"
	"fmt"
	"log"

	//"strconv"

	"yukbayar-rpll-be/db"
	"yukbayar-rpll-be/helpers/response"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context) {
	db := db.Connect()
	defer db.Close()

	id := c.Param("id")

	query := "SELECT * FROM pengguna WHERE ID =" + id
	//query := "SELECT * FROM pengguna"

	// if nama != "" {
	// 	query += "WHERE nama='" + nama + "'"
	// }

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var profile Pengguna
	var profiles []Pengguna

	for rows.Next() {
		if err := rows.Scan(&profile.ID, &profile.Email, &profile.Nama, &profile.NoTelpon, &profile.Password, &profile.TglLahir, &profile.Gender, &profile.SaldoYukPay, &profile.TipePengguna); err != nil {
			log.Fatal(err.Error())
		} else {
			profiles = append(profiles, profile)
		}
	}

	fmt.Println(profiles)

	var res response.Response
	if len(profiles) > 0 {
		res.Message = "Get User Data Success"
		res.Data = profiles
		response.SendSuccessResponse(c, res)
	} else {
		res.Message = "Get User Data Failed"
		response.SendErrorResponse(c, res)
	}

}

func UpdateUser(c *gin.Context) {

	db := db.Connect()
	defer db.Close()

	nama := c.PostForm("nama")
	pass := c.PostForm("password")
	noTelp := c.PostForm("noTelpon")
	id := c.Param("id")

	var profile Pengguna
	err := db.QueryRow("SELECT nama, noTelpon, password FROM pengguna WHERE id="+id).Scan(&profile.Nama, &profile.NoTelpon, &profile.Password)

	if err != nil {
		log.Fatal(err)
	}

	if err == sql.ErrNoRows {
		response.SendErrorResponse(c, response.Response{
			Message: "User ID not found",
		})
		return
	}

	if nama == "" {
		nama = profile.Nama
	}

	if noTelp == "" {
		noTelp = profile.NoTelpon
	}

	if pass == "" {
		pass = profile.Password
	}

	_, errQuery := db.Exec("UPDATE pengguna SET nama = ?, password = ?, noTelpon = ? WHERE id = ?",
		nama,
		pass,
		noTelp,
		id,
	)

	var res response.Response
	if errQuery == nil {
		res.Message = "Update Profile Success"
		response.SendSuccessResponse(c, res)
	} else {
		res.Message = "Update Profile Failed"
		response.SendErrorResponse(c, res)
	}
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
