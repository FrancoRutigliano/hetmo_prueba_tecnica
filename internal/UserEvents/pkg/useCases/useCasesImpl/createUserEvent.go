package userEventsUseCasesImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (u *UserEvents) CreateUserEvents(userId, eventId string) httpresponse.ApiResponse {

	err := u.UserEvents.Impl.CreateUserEvent(userId, eventId, u.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusCreated, "create", nil)
}
