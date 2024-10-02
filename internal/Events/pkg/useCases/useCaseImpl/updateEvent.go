package eventsUseCaseImpl

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"hetmo_prueba_tecnica/pkg/utils"
	"net/http"
	"time"
)

func (e *Events) UpdateEvent(userId, eventId string, payload eventsDto.EventResponseDTO) httpresponse.ApiResponse {
	var dto = eventsDto.EventResponseDTOUpdate{}
	var unixTime int64

	// Actualizar solo los campos presentes en el payload
	if payload.Title != "" {
		dto.Title = payload.Title
	}
	if payload.ShortDescription != "" {
		dto.ShortDescription = payload.ShortDescription
	}
	if payload.LongDescription != "" {
		dto.ShortDescription = payload.ShortDescription
	}
	if unixTime != 0 {
		unixTime, err := utils.ParseDateToUnix(payload.Date)
		if err != nil {
			return *httpresponse.NewApiError(http.StatusBadRequest, "invalid to parse date", nil)
		}
		dto.Date = unixTime
	}
	if payload.Location != "" {
		dto.Location = payload.Location
	}
	dto.Organizer = userId // Asigna el organizador actual si es necesario
	dto.IsPublished = payload.IsPublished
	dto.UpdatedAt = time.Now()

	data, err := e.EventsRepository.Impl.UpdateEvent(eventId, dto, e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "event edit succesfully", data)
}
