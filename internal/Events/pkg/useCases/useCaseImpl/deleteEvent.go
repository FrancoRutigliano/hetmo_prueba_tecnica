package eventsUseCaseImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (e *Events) DeleteEvent(id string) httpresponse.ApiResponse {
	err := e.EventsRepository.Impl.DeleteEvent(id, e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusAccepted, "event deleted succesfully", nil)
}
