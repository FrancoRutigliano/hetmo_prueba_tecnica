package userEventsUseCasesImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (u *UserEvents) CreateUserEvents(userId, eventId string) httpresponse.ApiResponse {

	isActive, err := u.UserEvents.Impl.IsEventActive(eventId, u.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusBadRequest, err.Error(), nil)
	}
	if !isActive {
		return *httpresponse.NewApiError(http.StatusBadRequest, "The event is not active", nil)
	}

	err = u.UserEvents.Impl.CreateUserEvent(userId, eventId, u.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusCreated, "successfully registered user", nil)
}
