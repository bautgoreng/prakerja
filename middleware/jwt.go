package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJwt(name string, id int) (string, error) {
	claims := &jwtCustomClaims{
		name,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("PRIVATE_KEY_JWT")))
	if err != nil {
		return "", err
	}
	return t, nil
}
