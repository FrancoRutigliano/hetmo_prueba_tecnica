package eventsUseCaseImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (e *Events) GetEvents() httpresponse.ApiResponse {
	data, err := e.EventsRepository.Impl.GetAllEvents(e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "oops somenthing went wrong", nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "all events", data)
}
