package usecaseimpl

import (
	"errors"
	"fmt"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	utilsAuth "hetmo_prueba_tecnica/pkg/auth"
	"time"

	"github.com/google/uuid"
)

func (a *Auth) Register(payload authDto.AuthRegisterPayload) (string, error) {
	err := a.Repository.Impl.FindByEmail(payload.Email, a.Db)
	if err != nil {
		return "", errors.New("email already exist")
	}

	userID := uuid.New().String()

	hashPass, err := utilsAuth.HashPassword(payload.Password)
	if err != nil {
		return "", errors.New("error to hash password")
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
		return "", fmt.Errorf("error to register user -> %s ", err.Error())
	}

	return "user created", nil
}
