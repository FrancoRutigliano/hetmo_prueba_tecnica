package authController

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	"hetmo_prueba_tecnica/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) Login(c *fiber.Ctx) error {
	// sacar el body de la request
	var payload authDto.AuthLoginRequest

	// Pasar la referencia a `&payload` en lugar de `payload`
	response := utils.ValidPayload(c, &payload)
	if response.StatusCode != 0 {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}

	response = a.handler.AuthCase.Login(payload)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "true"})
}
