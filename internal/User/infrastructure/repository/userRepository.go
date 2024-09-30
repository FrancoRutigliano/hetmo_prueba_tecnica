package infraUserRepository

import (
	"fmt"
	userDto "hetmo_prueba_tecnica/internal/User/pkg/domain/dto"
	userRepository "hetmo_prueba_tecnica/internal/User/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type SqlxUserRepository struct {
	Impl userRepository.Repository
}

type ImplrepositoryUser struct {
}

func (s *SqlxUserRepository) New() {
	s.Impl = &ImplrepositoryUser{}
}

func (s *ImplrepositoryUser) FindUserById(id string, db *sqlx.DB) (*userDto.UserInfoResponse, error) {
	query := `SELECT u.id, u.name, u.email, u.role FROM "public".users AS u WHERE id=$1`

	var payload userDto.UserInfoResponse

	if err := db.Get(&payload, query, id); err != nil {
		return nil, fmt.Errorf("error finding user with ID %s", id)
	}

	return &payload, nil
}
