package usecaseimpl

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	utilsAuth "hetmo_prueba_tecnica/pkg/auth"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (a *Auth) Register(payload authDto.AuthRegisterPayload) httpresponse.ApiResponse {
	err := a.Repository.Impl.FindByEmail(payload.Email, a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusBadRequest, "email already exists")
	}

	userID := uuid.New().String()

	hashPass, err := utilsAuth.HashPassword(payload.Password)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "oops somenthing went wrong")
	}

	var dto = authDto.AuthRegisterRequest{
		Id:       userID,
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashPass,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = a.Repository.Impl.RegisterUser(dto, a.Db)
	if err != nil {
		return *httpresponse.NewApiError(http.StatusInternalServerError, "oops somenthing went wrong")
	}

	return *httpresponse.NewApiError(http.StatusCreated, "user created succesfully")
}
