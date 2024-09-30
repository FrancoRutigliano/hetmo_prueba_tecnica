package authJwt

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Create(secret []byte, id string, role bool) (string, error) {
	expiration := os.Getenv("JWT_EXP")

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    id,
		"role":      role,
		"expiredAt": expiration,
	})

	token, err := t.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(t string) (*jwt.Token, error) {
	tokenStr := strings.TrimPrefix(t, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetRole(token *jwt.Token) (string, error) {

	if !token.Valid {
		return "", fmt.Errorf("invalid token on get role")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	roleBool, ok := claims["role"].(bool)
	if !ok {
		return "", fmt.Errorf("role not found in claims")
	}

	role := strconv.FormatBool(roleBool)

	return role, nil
}
