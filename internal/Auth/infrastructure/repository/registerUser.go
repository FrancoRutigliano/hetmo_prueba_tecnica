package infraAuthRepository

import (
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"

	"github.com/jmoiron/sqlx"
)

func (s *ImplrepositoryAuth) RegisterUser(payload authDto.AuthRegisterRequest, db *sqlx.DB) error {
	query := `INSERT INTO  "public".users (id, name, email, password, role, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := db.Exec(query, payload.Id, payload.Name, payload.Email, payload.Password, payload.Role, payload.CreatedAt, payload.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
