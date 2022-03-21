package Profile

import (
	//"errors"
	"log"

	//"strconv"

	"github.com/gin-gonic/gin"
	repo "github.com/yukbayar-backend/db_handler"
	"github.com/yukbayar-backend/model"
)

func GetUserData(c *gin.Context) {
	db := repo.Connect()
	defer db.Close()

	nama := c.Param("nama")

	query := "SELECT * FROM pengguna WHERE nama ='" + nama + "'"
	//query := "SELECT * FROM pengguna"

	// if nama != "" {
	// 	query += "WHERE nama='" + nama + "'"
	// }

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var profile model.Pengguna
	var profiles []model.Pengguna

	for rows.Next() {
		if err := rows.Scan(&profile.ID, &profile.Email, &profile.Nama, &profile.NoTelpon, &profile.Password, &profile.TglLahir, &profile.Gender, &profile.SaldoYukPay, &profile.TipePengguna); err != nil {
			log.Fatal(err.Error())
		} else {
			profiles = append(profiles, profile)
		}
	}

	var response model.PenggunaResponse
	if err == nil {
		response.Message = "Get User Data Success"
		response.Data = profiles
		sendUserSuccessresponse(c, response)
	} else {
		response.Message = "Get User Data Failed"
		sendUserErrorresponse(c, response)
	}

}

func UpdateUser(c *gin.Context) {

	db := repo.Connect()
	defer db.Close()

	nama := c.PostForm("nama")
	pass := c.PostForm("password")
	noTelp := c.PostForm("noTelpon")
	id := c.Param("id")

	rows, _ := db.Query("SELECT * FROM pengguna WHERE id='" + id + "'")
	var profile model.Pengguna
	for rows.Next() {
		if err := rows.Scan(&profile.ID, &profile.Email, &profile.Nama, &profile.NoTelpon, &profile.Password, &profile.TglLahir, &profile.Gender, &profile.SaldoYukPay, &profile.TipePengguna); err != nil {
			log.Fatal(err.Error())
		}
	}

	if nama == "" {
		nama = profile.Nama
	}

	if noTelp == "" {
		noTelp = profile.NoTelpon
	}

	_, errQuery := db.Exec("UPDATE pengguna SET nama = ?, password = ?, noTelpon = ? WHERE id = ?",
		nama,
		pass,
		noTelp,
		id,
	)

	var response model.PenggunaResponse
	if errQuery == nil {
		response.Message = "Update Profile Success"
		sendUserSuccessresponse(c, response)
	} else {
		response.Message = "Update Profile Success"
		sendUserErrorresponse(c, response)
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
