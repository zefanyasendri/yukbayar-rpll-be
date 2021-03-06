package main

import (
	"fmt"
	"time"

	"yukbayar-rpll-be/db"
	"yukbayar-rpll-be/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.Connect()
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

	fmt.Println(time.Now())

	routes.Routes(router, db)

	router.Run(":8000")
	fmt.Println("Connected to port 8000")
}
