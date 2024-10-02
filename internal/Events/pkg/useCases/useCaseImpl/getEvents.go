package eventsUseCaseImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (e *Events) GetEvents() httpresponse.ApiResponse {
	data, err := e.EventsRepository.Impl.GetAllEvents(e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "all events", data)
}
