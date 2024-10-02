package eventsController

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (e *Events) GetEvents(c *fiber.Ctx) error {
	// query param state and date
	title := c.Query("title", "")
	isPublished := c.QueryBool("is_pusblished", true)
	dateStr := c.QueryInt("date", 0)

	var dto = eventsDto.GetEventsRequest{
		Title:       title,
		IsPublished: isPublished,
		Date:        int64(dateStr),
	}

	response := e.handler.EventsCase.GetEvents(dto)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
}
