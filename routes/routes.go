package routes

import (
	"yukbayar-rpll-be/models/pengguna"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	users := route.Group("/users")
	{
		users.GET("/:id", pengguna.GetUserData)
		users.PUT("/:id", pengguna.UpdateUser)
	}
}
