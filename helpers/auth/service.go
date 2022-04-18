package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(email string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type service struct {
	secretKey string
	issure    string
}

func NewService() Service {
	return &service{
		secretKey: getSecretKey(),
		issure:    "YukBayar",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "YukBayar"
	}
	return secret
}

func (s *service) GenerateToken(email string) (string, error) {
	claims := &Claims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))

	return t, err
}

func (s *service) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}
