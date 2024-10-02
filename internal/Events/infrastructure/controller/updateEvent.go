package eventsController

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	authJwt "hetmo_prueba_tecnica/pkg/jwt"
	"hetmo_prueba_tecnica/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (e *Events) UpdateEvent(c *fiber.Ctx) error { // arreglar que esta sobreescribiendo
	var payload eventsDto.EventResponseDTO
	id := c.Query("id", "")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "user id required", "details": "false"})
	}

	response := utils.ValidPayload(c, &payload)
	if response.StatusCode != 0 {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}

	userId, err := authJwt.GetID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "details": "false"})
	}

	response = e.handler.EventsCase.UpdateEvent(userId, id, payload)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "true"})
}
