package routes

import (
	"yukbayar-rpll-be/modules/pengguna"
	"yukbayar-rpll-be/modules/walletTopUp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine, db *gorm.DB) {
	penggunaRepository := pengguna.NewRepository(db)
	penggunaService := pengguna.NewService(penggunaRepository)
	penggunaController := pengguna.NewController(penggunaService)
	topupRepository := walletTopUp.NewRepository(db)
	topupService := walletTopUp.NewService(topupRepository)
	topupController := walletTopUp.NewController(topupService)

	root := route.Group("/")
	{
		root.POST("/register", penggunaController.Register)
		root.GET("/login", penggunaController.Login)
		root.GET("/check", penggunaController.CheckLogin)
	}

	users := route.Group("/users")
	{
		users.GET("/:id", penggunaController.GetUserData)
		users.PUT("/:id", penggunaController.UpdateUser)

	}

	topups := route.Group("/topups")
	{
		topups.POST("/topup", topupController.InsertNewTopup)
		topups.PUT("/:id", topupController.UpdateSaldoUser)
	}
}
