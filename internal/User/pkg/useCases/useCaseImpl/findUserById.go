package userUseCaseImpl

import (
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
)

func (u *User) FindUserById(id string) httpresponse.ApiResponse {
	user, err := u.Repository.Impl.FindUserById(id, u.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusBadRequest, err.Error(), nil)
	}
	return *httpresponse.NewApiError(http.StatusOK, "user found!", user)
}
