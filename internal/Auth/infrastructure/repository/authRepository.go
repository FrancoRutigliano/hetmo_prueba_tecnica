package infraAuthRepository

import (
	"database/sql"
	"fmt"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	authRepository "hetmo_prueba_tecnica/internal/Auth/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type SqlxAuthRepository struct {
	Impl authRepository.Repository
}

type ImplrepositoryAuth struct {
}

func (s *SqlxAuthRepository) New() {
	s.Impl = &ImplrepositoryAuth{}
}

func (s *ImplrepositoryAuth) FindByEmail(emailParam string, db *sqlx.DB) error {
	query := `SELECT u.email FROM "public".users AS u WHERE email=$1 LIMIT 1;`

	var email string

	if err := db.Get(&email, query, emailParam); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return fmt.Errorf("error executing the query")
	}
	return fmt.Errorf("email: %s already exist", emailParam)
}

func (s *ImplrepositoryAuth) RegisterUser(payload authDto.AuthRegisterRequest, db *sqlx.DB) error {
	query := `INSERT INTO  "public".users (id, name, email, password, role, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7);`

	_, err := db.Exec(query, payload.Id, payload.Name, payload.Email, payload.Password, payload.Role, payload.CreatedAt, payload.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *ImplrepositoryAuth) GetUser(payload authDto.AuthLoginRequest, db *sqlx.DB) (authDto.AuthLoginResponse, error) {
	query := `SELECT u.id, u.email, u.password FROM "public".users AS u WHERE email=$1 LIMIT 1;`

	var user authDto.AuthLoginResponse

	if err := db.Get(&user, query, payload.Email); err != nil {
		if err == sql.ErrNoRows {
			return authDto.AuthLoginResponse{}, fmt.Errorf("user not found for email: %s", payload.Email)
		}
		return authDto.AuthLoginResponse{}, fmt.Errorf("error executing query: %v", err)
	}
	return user, nil
}
