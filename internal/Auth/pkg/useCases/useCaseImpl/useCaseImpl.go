package usecaseimpl

import (
	infraAuthRepository "hetmo_prueba_tecnica/internal/Auth/infrastructure/repository"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"

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
