package infraAuthRepository

import (
	authRepository "hetmo_prueba_tecnica/internal/Auth/pkg/domain/repository"

	"github.com/jmoiron/sqlx"
)

type SqlxAuthRepository struct {
	repo authRepository.Repository
}

type ImplrepositoryAuth struct {
}

func (s *SqlxAuthRepository) New() {
	s.repo = &ImplrepositoryAuth{}
}

func (s *ImplrepositoryAuth) FindByEmail(email string, db *sqlx.DB) (string, error) {
	return "", nil
}
