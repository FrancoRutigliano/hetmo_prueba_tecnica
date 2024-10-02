package eventsController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (e *Events) GetCompleteEvents(c *fiber.Ctx) error {
	title := c.Query("title", "")

	response := e.handler.EventsCase.GetCompletedEvents(title)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "true"})
}
