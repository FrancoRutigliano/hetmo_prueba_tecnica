package authController

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	"hetmo_prueba_tecnica/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) Login(c *fiber.Ctx) error {
	// sacar el body de la request
	var payload authDto.AuthLoginRequest

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload"})
	}

	if err := utils.Validate.Struct(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Invalid entity"})
	}

	response, err := a.handler.AuthCase.Login(payload)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": response, "details": "false"})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message ": "todo ok", "details": "true"})
}
