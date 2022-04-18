package routes

import (
	"yukbayar-rpll-be/modules/pengguna"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(route *gin.Engine, db *gorm.DB) {
	penggunaRepository := pengguna.NewRepository(db)
	penggunaService := pengguna.NewService(penggunaRepository)
	penggunaController := pengguna.NewController(penggunaService)

	root := route.Group("/")
	{
		root.POST("/register", penggunaController.Register)
		root.GET("/login", penggunaController.Login)
	}

	users := route.Group("/users")
	{
		users.GET("/:id", penggunaController.GetUserData)
		users.PUT("/:id", penggunaController.UpdateUser)
	}
}
