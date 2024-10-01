package infraAuthRepository

import (
	authRepository "hetmo_prueba_tecnica/internal/Auth/pkg/domain/repository"
)

type SqlxAuthRepository struct {
	Impl authRepository.Repository
}

type ImplrepositoryAuth struct {
}

func (s *SqlxAuthRepository) New() {
	s.Impl = &ImplrepositoryAuth{}
}
