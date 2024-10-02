package userEventsUseCasesImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (u *UserEvents) DeleteUserEvents(userId, eventId string) httpresponse.ApiResponse {
	err := u.UserEvents.Impl.DeleteUserEvent(userId, eventId, u.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "oops somenthing went wrong", nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "succesfully user deleted", nil)
}
