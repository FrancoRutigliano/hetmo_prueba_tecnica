package eventsController

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (e *Events) GetEvents(c *fiber.Ctx) error {
	// query param state and date
	isPublishedStr := c.Query("is_pusblished", "")
	var isPublished bool
	isPublished = strings.EqualFold(isPublishedStr, "true")
	dateStr := c.Query("date", "")
	date, err := time.Parse(time.RFC3339, dateStr) // Asegúrate de que el formato sea correcto
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid date format"})
	}

	var dto = eventsDto.GetEventsRequest{
		IsPublished: isPublished,
		Date:        date,
	}

	response := e.handler.EventsCase.GetEvents(dto)
	if response.StatusCode != http.StatusOK {
		return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}
	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "details": "false"})
}
