package eventsController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (e *Events) GetEvents(c *fiber.Ctx) error {

	response := e.handler.EventsCase.GetEvents()
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "false"})
}
