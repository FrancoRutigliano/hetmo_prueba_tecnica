package usecaseimpl

import (
	"errors"
	"fmt"
	infraAuthRepository "hetmo_prueba_tecnica/internal/Auth/infrastructure/repository"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	utilsAuth "hetmo_prueba_tecnica/pkg/auth"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IAuthUseCase interface {
	Login(authDto.AuthLoginRequest) (string, error)
	Register(authDto.AuthRegisterPayload) (string, error)
}

type Auth struct {
	Repository infraAuthRepository.SqlxAuthRepository
	Db         *sqlx.DB
}

func (a *Auth) Login(payload authDto.AuthLoginRequest) (string, error) {

	return "Login", nil
}

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
