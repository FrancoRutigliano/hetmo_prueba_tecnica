package eventsController

import "github.com/gofiber/fiber/v2"

func (e *Events) CreateEvent(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "gol", "details": "true"})
}
