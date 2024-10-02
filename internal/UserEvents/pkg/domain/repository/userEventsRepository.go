package userEventsRepository

import (
	userEventsDto "hetmo_prueba_tecnica/internal/UserEvents/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateUserEvent(string, string, *sqlx.DB) error
	GetUserEvent(string, string, *sqlx.DB) ([]userEventsDto.UserEventsListDto, error)
	DeleteUserEvent(string, string, *sqlx.DB) error
	IsEventActive(string, *sqlx.DB) (bool, error)
}
