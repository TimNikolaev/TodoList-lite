package jwt

import (
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
	id := token.Claims.(jwt.MapClaims)["id"]
	return id.(int), err
}
