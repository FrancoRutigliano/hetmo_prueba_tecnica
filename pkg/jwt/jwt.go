package authJwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Create(secret []byte, id string) (string, error) {
	expiration := os.Getenv("JWT_EXP")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    id,
		"expiredAt": expiration,
	})

	token, err := t.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
