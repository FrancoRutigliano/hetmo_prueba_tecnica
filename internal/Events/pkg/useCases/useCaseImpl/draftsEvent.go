package eventsUseCaseImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (e *Events) GetDraftEvents(title string) httpresponse.ApiResponse {
	data, err := e.EventsRepository.Impl.GetDraftEvents(title, e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "oops somenthing went wrong", nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "events actives founded", data)
}
