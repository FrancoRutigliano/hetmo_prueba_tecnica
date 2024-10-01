package authController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) ChangePassword(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "change pass"})
}
