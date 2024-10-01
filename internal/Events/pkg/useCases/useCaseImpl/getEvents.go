package eventsUseCaseImpl

import (
	eventsDto "hetmo_prueba_tecnica/internal/Events/pkg/domain/dto"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (e *Events) GetEvents(dto eventsDto.GetEventsRequest) httpresponse.ApiResponse {
	data, err := e.EventsRepository.Impl.GetAllEvents(&dto.IsPublished, &dto.Date, e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "all events", data)
}
