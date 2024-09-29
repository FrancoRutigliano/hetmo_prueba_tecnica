package usecaseimpl

import (
	infraAuthRepository "hetmo_prueba_tecnica/internal/Auth/infrastructure/repository"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"

	"github.com/jmoiron/sqlx"
)

type IAuthUseCase interface {
	Login(authDto.AuthLoginRequest) httpresponse.ApiResponse
	Register(authDto.AuthRegisterPayload) httpresponse.ApiResponse
}

type Auth struct {
	Repository infraAuthRepository.SqlxAuthRepository
	Db         *sqlx.DB
}
