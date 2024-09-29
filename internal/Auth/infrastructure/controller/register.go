package authController

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	"hetmo_prueba_tecnica/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) Register(c *fiber.Ctx) error {
	var payload authDto.AuthRegisterPayload

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request payload", "details": "false"})
	}

	if err := utils.Validate.Struct(&payload); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Invalid entity", "details": "false"})
	}

	response, err := a.handler.AuthCase.Register(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "details": "false"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": response, "details": "true"})
}
