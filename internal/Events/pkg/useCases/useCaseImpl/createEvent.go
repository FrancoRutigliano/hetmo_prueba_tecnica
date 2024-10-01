package eventsUseCaseImpl

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (e *Events) CreateEvent(event eventsDto.EventCreateDTORequest, userId string, unixTime int64) httpresponse.ApiResponse {

	eventId := uuid.New().String()

	var dto = eventsDto.EventCreateDTO{
		Id:               eventId,
		Title:            event.Title,
		ShortDescription: event.ShortDescription,
		LongDescription:  event.ShortDescription,
		Date:             unixTime,
		Organizer:        userId,
		Location:         event.Location,
		IsPublished:      false,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := e.EventsRepository.Impl.CreateEvent(dto, e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "event created succesfully", dto)
}
