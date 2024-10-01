package middleware

import (
	authJwt "hetmo_prueba_tecnica/pkg/jwt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UserMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "no token provided", "details": "false"})
	}

	tokenStr, err := authJwt.ValidateToken(token)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"message": err.Error(), "details": "false"})
	}

	_, err = authJwt.GetRole(tokenStr)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": err.Error(), "details": "false"})
	}
	return c.Next()
}
