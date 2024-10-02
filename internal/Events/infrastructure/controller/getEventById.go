package eventsController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (e *Events) GetEventById(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "user id required", "details": "false"})
	}

	response := e.handler.EventsCase.GetEventByID(id)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "true"})
}
