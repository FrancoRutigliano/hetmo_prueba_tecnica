package authController

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	"hetmo_prueba_tecnica/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) Register(c *fiber.Ctx) error {
	var payload authDto.AuthRegisterPayload

	response := utils.ValidPayload(c, &payload)
	if response.StatusCode != 0 {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}

	response = a.handler.AuthCase.Register(payload)
	if response.StatusCode != http.StatusCreated {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "true"})
}
