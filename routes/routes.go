package routes

import (
	"yukbayar-rpll-be/controllers"
	"yukbayar-rpll-be/repositories"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine, db *gorm.DB) {
	penggunaRepository := repositories.NewPenggunaRepository(db)
	penggunaService := services.NewPenggunaService(penggunaRepository)
	penggunaController := controllers.NewPenggunaController(penggunaService)

	varianRepository := repositories.NewVarianRepository(db)
	varianService := services.NewVarianService(varianRepository)
	varianController := controllers.NewVarianController(varianService)

	kategoriLayananRepository := repositories.NewKategoriLayananRepository(db)
	kategoriLayananService := services.NewKategoriLayananService(kategoriLayananRepository)
	kategoriLayananController := controllers.NewKategoriLayananController(kategoriLayananService)

	root := route.Group("/")
	{
		root.POST("/register", penggunaController.Register)
		root.GET("/login", penggunaController.Login)
		root.GET("/varian/:id", varianController.GetVarian)
		root.GET("/kategori/:id", kategoriLayananController.GetKategori)
	}

	users := route.Group("/users")
	{
		users.GET("/:id", penggunaController.GetUserData)
		users.PUT("/:id", penggunaController.UpdateUser)
	}
}
