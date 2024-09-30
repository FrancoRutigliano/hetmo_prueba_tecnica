package authUsecaseimpl

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	utilsAuth "hetmo_prueba_tecnica/pkg/auth"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	authJwt "hetmo_prueba_tecnica/pkg/jwt"
	"net/http"
	"os"
)

func (a *Auth) Login(payload authDto.AuthLoginRequest) httpresponse.ApiResponse {

	user, err := a.Repository.Impl.GetUser(payload, a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusBadRequest, "email doesn't exists", nil)
	}

	if !utilsAuth.ComparePasswords(user.Password, []byte(payload.Password)) {
		return *httpresponse.NewApiError(http.StatusBadRequest, "incorrect password", nil)
	}

	secret := []byte(os.Getenv("SECRET_JWT"))

	token, err := authJwt.Create(secret, user.Id, user.Role)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "oops somenthing went wrong", nil)
	}

	return *httpresponse.NewApiError(http.StatusOK, token, nil)
}
