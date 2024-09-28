package authRepository

import "github.com/jmoiron/sqlx"

type Repository interface {
	FindByEmail(string, *sqlx.DB) (string, error)
}
