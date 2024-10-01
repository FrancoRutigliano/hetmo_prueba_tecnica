package eventsUseCaseImpl

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
	"time"
)

func (e *Events) UpdateEvent(userId, eventId string, payload eventsDto.EventResponseDTO) httpresponse.ApiResponse {

	var dto = eventsDto.EventResponseDTO{
		Title:            payload.Title,
		ShortDescription: payload.ShortDescription,
		Date:             payload.Date,
		Organizer:        userId,
		Location:         payload.Location,
		IsPublished:      payload.IsPublished,
		UpdatedAt:        time.Now(),
	}

	data, err := e.EventsRepository.Impl.UpdateEvent(eventId, dto, e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "event edit succesfully", data)
}
