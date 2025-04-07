package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenTTL = 24 * time.Hour
)

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Generate(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(tokenTTL).Unix(),
	})

	return token.SignedString([]byte(j.Secret))
}

func (j *JWT) Parse(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (any, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	floatID, ok := claims["id"].(float64)
	if !ok {
		return 0, fmt.Errorf("id is not a number")
	}

	return int(floatID), nil

}
