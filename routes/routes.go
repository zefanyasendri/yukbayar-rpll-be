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

	layananRepository := repositories.NewLayananRepository(db)
	layananService := services.NewLayananService(layananRepository)
	layananController := controllers.NewLayananController(layananService)

	transaksiRepository := repositories.NewTransaksiRepository(db)
	transaksiService := services.NewTransaksiService(transaksiRepository)
	transaksiController := controllers.NewTransaksiController(transaksiService)

	topUpRepository := repositories.NewTopUpRepository(db)
	topUpService := services.NewTopUpService(topUpRepository)
	topUpController := controllers.NewTopUpController(topUpService)

	diskonRepository := repositories.NewDiskonRepository(db)
	diskonService := services.NewDiskonService(diskonRepository)
	diskonController := controllers.NewDiskonController(diskonService)

	mitraRepository := repositories.NewMitraRepository(db)
	mitraService := services.NewMitraService(mitraRepository)
	mitraController := controllers.NewMitraController(mitraService)

	root := route.Group("/")
	{
		root.POST("/register", penggunaController.Register)
		root.POST("/login", penggunaController.Login)
		root.GET("/varian/:id", varianController.GetVarian)
		root.GET("/kategori/:id", kategoriLayananController.GetKategori)
		root.GET("/layanan", layananController.GetLayanan)
		root.GET("/mitra", mitraController.GetAllMitra)
		root.GET("/pendapatan", transaksiController.GetTotalHarga)
		root.GET("/topup", topUpController.GetAll)
		root.GET("/topup/:id", topUpController.GetByTopUpID)
	}

	users := route.Group("/users")
	{
		users.GET("/:id", penggunaController.GetUserData)
		users.GET("/", penggunaController.GetAllData)
		users.PUT("/:id", penggunaController.UpdateUser)
		users.PUT("/:id/saldo", penggunaController.UpdateSaldo)

	}

	saldo := route.Group("/saldo")
	{
		saldo.GET("/:idPengguna", topUpController.GetBySaldoByID)
	}

	topups := route.Group("/topups")
	{
		topups.POST("/topup", topUpController.InsertNewTopup)
		topups.PUT("/:idPengguna/:id", topUpController.UpdateSaldoUser)
		topups.GET("/:id", topUpController.GetByPenggunaID)
	}

	diskons := route.Group("/diskon")
	{
		diskons.GET("/", diskonController.GetAll)
	}

	transaksis := route.Group("/transaksi")
	{
		transaksis.GET("/", transaksiController.GetTransaksi)
		transaksis.POST("/", transaksiController.CreateTransaksi)
		transaksis.GET("/:id", transaksiController.GetTransaksiById)
		//transaksis.PUT("/status", transaksiController.UpdateStatusTransaksi)
	}

	mitras := route.Group("/mitras")
	{
		mitras.POST("/mitra", mitraController.InsertMitra)
	}
}
