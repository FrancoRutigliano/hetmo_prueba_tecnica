package utils

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func ValidPayload(c *fiber.Ctx, payload interface{}) httpresponse.ApiResponse {

	if err := c.BodyParser(payload); err != nil {
		return *httpresponse.NewApiError(http.StatusBadRequest, "invalid payload request")
	}

	if err := Validate.Struct(payload); err != nil {
		return *httpresponse.NewApiError(http.StatusUnprocessableEntity, "invalid entity")
	}

	return httpresponse.ApiResponse{}
}
