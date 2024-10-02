package userEventsUseCasesImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
)

func (u *UserEvents) GetUserEvents() httpresponse.ApiResponse {
	return *httpresponse.NewApiError(200, "get", nil)
}
