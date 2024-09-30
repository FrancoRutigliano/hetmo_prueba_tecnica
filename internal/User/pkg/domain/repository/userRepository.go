package userRepository

import (
	userDto "hetmo_prueba_tecnica/internal/User/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	FindUserById(string, *sqlx.DB) (*userDto.UserInfoResponse, error)
}
