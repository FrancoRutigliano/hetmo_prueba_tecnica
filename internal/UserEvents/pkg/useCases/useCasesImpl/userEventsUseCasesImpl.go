package userEventsUseCasesImpl

import (
	infraUserEventsRepository "hetmo_prueba_tecnica/internal/UserEvents/infrastructure/repository"
	httpresponse "hetmo_prueba_tecnica/pkg/httpResponse"

	"github.com/jmoiron/sqlx"
)

type IUserEventsUseCase interface {
	CreateUserEvents(string, string) httpresponse.ApiResponse
	GetUserEvents(string, string) httpresponse.ApiResponse
	DeleteUserEvents(string, string) httpresponse.ApiResponse
}

type UserEvents struct {
	UserEvents infraUserEventsRepository.SqlxUserEventsRepository
	Db         *sqlx.DB
}
