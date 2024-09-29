package authRepository

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindByEmail(string, *sqlx.DB) error
	RegisterUser(authDto.AuthRegisterRequest, *sqlx.DB) error
	GetUser(authDto.AuthLoginRequest, *sqlx.DB) (authDto.AuthLoginRequest, error)
}
