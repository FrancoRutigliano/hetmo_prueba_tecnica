package eventsUseCaseImpl

import (
	"fmt"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (e *Events) GetEventByID(id string) httpresponse.ApiResponse {
	data, err := e.EventsRepository.Impl.GetEventById(id, e.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	msg := fmt.Sprintf("id: %s event founded!", id)
	return *httpresponse.NewApiError(http.StatusOK, msg, data)
}
