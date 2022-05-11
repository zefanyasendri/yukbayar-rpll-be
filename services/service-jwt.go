package services

import (
	"fmt"
	"time"
	"yukbayar-rpll-be/models"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(pengguna models.Pengguna) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Pengguna models.Pengguna `json:"pengguna"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "YukBayar-secret",
		issuer:    "YukBayar",
	}
}

func (s *jwtService) GenerateToken(pengguna models.Pengguna) string {
	claims := &jwtCustomClaims{
		pengguna,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}
