package controllers

import (
	"log"
	"yukbayar-rpll-be/helpers"
	"yukbayar-rpll-be/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helpers.SendErrorResponse(c, helpers.Response{
				Message: "Token not found",
				Data:    nil,
			})
		}

		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[pengguna]: ", claims["pengguna"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			helpers.SendErrorResponse(c, helpers.Response{
				Message: "Token not valid",
				Data:    err.Error(),
			})
		}
	}
}
