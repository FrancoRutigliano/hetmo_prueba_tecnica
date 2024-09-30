package authUseCase

import (
	infraAuthRepository "hetmo_prueba_tecnica/internal/Auth/infrastructure/repository"
	authUsecaseimpl "hetmo_prueba_tecnica/internal/Auth/pkg/useCases/useCaseImpl"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/data"
	"log"
)

type AuthImpl struct {
	AuthCase authUsecaseimpl.IAuthUseCase
}

func (a *AuthImpl) New() {
	var repository infraAuthRepository.SqlxAuthRepository

	repository.New()

	db, err := data.GetConnection()
	if err != nil {
		log.Println("error to connect db  --> ", err)
	}

	a.AuthCase = &authUsecaseimpl.Auth{
		Repository: repository,
		Db:         db,
	}

}
