package auth

import "github.com/golang-jwt/jwt"

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

func (c *controller) GenerateToken(email string) (string, error) {
	token, err := c.service.GenerateToken(email)
	return token, err
}

func (c *controller) ValidateToken(token string) (*jwt.Token, error) {
	return nil, nil
}
