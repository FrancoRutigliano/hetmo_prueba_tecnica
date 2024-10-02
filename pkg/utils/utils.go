package utils

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func ValidPayload(c *fiber.Ctx, payload interface{}) httpresponse.ApiResponse {

	if err := c.BodyParser(payload); err != nil {
		return *httpresponse.NewApiError(http.StatusBadRequest, err.Error(), nil)
	}

	if err := Validate.Struct(payload); err != nil {
		return *httpresponse.NewApiError(http.StatusUnprocessableEntity, "invalid entity", nil)
	}

	return httpresponse.ApiResponse{}
}

func ParseDateToUnix(dateStr string) (int64, error) {
	layout := "02/01/2006 15:04" // Formato DD/MM/YYYY HH:mm
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
