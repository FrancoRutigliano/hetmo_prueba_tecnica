package userEventsRepository

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateUserEvent(string, string, *sqlx.DB) error
	GetUserEvent(*sqlx.DB) error
	DeleteUserEvent(*sqlx.DB) error
}
