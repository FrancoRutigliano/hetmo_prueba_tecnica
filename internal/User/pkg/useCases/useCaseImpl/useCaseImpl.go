package userUseCaseImpl

import (
	infraUserRepository "hetmo_prueba_tecnica/internal/User/infrastructure/repository"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"

	"github.com/jmoiron/sqlx"
)

type IUserUseCase interface {
	FindUserById(string) httpresponse.ApiResponse
}

type User struct {
	Repository infraUserRepository.SqlxUserRepository
	Db         *sqlx.DB
}
