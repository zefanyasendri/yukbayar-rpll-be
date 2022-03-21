package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	profile "github.com/yukbayar-backend/profile"
)

func main() {
	fmt.Println("REST API Yuk Bayar")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	//router.GET("/user", profile.GetUserData)
	user := router.Group("/user")
	{
		user.GET("/:nama", profile.GetUserData)
		user.PUT("/:id", profile.UpdateUser)
	}

	router.Run(":8080")
	fmt.Println("Connected to port 8080")
}
