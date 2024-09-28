package usecaseimpl

import (
	"fmt"
	infraAuthRepository "hetmo_prueba_tecnica/internal/Auth/infrastructure/repository"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

type IAuthUseCase interface {
	Login(authDto.AuthLoginRequest) (string, error)
	Register(authDto.AuthRegisterRequest) (string, error)
}

type Auth struct {
	Repository infraAuthRepository.SqlxAuthRepository
	Db         *sqlx.DB
}

func (a *Auth) Login(payload authDto.AuthLoginRequest) (string, error) {

	return "Login", nil
}

func (a *Auth) Register(payload authDto.AuthRegisterRequest) (string, error) {

	response := fmt.Sprintf("Nombre: %s, Email: %s, Password: %s", payload.Name, payload.Email, payload.Password)

	return response, nil
}
