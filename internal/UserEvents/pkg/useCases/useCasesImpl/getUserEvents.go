package userEventsUseCasesImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (u *UserEvents) GetUserEvents(userId, eventId string) httpresponse.ApiResponse {

	data, err := u.UserEvents.Impl.GetUserEvent(userId, eventId, u.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "oops somenthing went wrong", nil)
	}

	return *httpresponse.NewApiError(http.StatusOK, "active events for this user", data)
}
