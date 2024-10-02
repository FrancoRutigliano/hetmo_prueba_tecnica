package eventsUseCaseImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (e *Events) GetPublishedEvents(title string) httpresponse.ApiResponse {
	data, err := e.EventsRepository.Impl.GetPublishedEvents(title, e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "events actives founded", data)
}
