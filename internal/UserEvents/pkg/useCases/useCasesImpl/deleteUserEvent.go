package userEventsUseCasesImpl

import httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"

func (u *UserEvents) DeleteUserEvents(userID, eventId string) httpresponse.ApiResponse {
	return *httpresponse.NewApiError(200, "delete", nil)
}
