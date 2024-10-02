package userEventsController

import (
	authJwt "hetmo_prueba_tecnica/pkg/jwt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (u *UserEvents) GeteUserEvent(c *fiber.Ctx) error {
	eventId := c.Query("id_event", "")

	userId, err := authJwt.GetID(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "details": "false"})
	}

	response := u.handler.UserEventsCase.GetUserEvents(userId, eventId)
	if response.StatusCode != http.StatusOK {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": response.Msg, "details": "false"})
	}

	return c.Status(response.StatusCode).JSON(fiber.Map{"message": response.Msg, "data": response.Data, "details": "true"})
}
