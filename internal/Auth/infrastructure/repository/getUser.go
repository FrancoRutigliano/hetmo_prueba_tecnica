package infraAuthRepository

import (
	"database/sql"
	"fmt"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (s *ImplrepositoryAuth) GetUser(payload authDto.AuthLoginRequest, db *sqlx.DB) (authDto.AuthLoginResponse, error) {
	query := `SELECT u.id, u.email, u.password, u.role FROM "public".users AS u WHERE email=$1 LIMIT 1;`

	var user authDto.AuthLoginResponse

	if err := db.Get(&user, query, payload.Email); err != nil {
		if err == sql.ErrNoRows {
			return authDto.AuthLoginResponse{}, fmt.Errorf("user not found for email: %s", payload.Email)
		}
		return authDto.AuthLoginResponse{}, fmt.Errorf("error executing query: %v", err)
	}
	return user, nil
}
