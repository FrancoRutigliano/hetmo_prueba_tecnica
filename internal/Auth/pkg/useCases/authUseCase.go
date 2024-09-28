package authUseCase

import (
	infraAuthRepository "hetmo_prueba_tecnica/internal/Auth/infrastructure/repository"
	usecaseimpl "hetmo_prueba_tecnica/internal/Auth/pkg/useCases/useCaseImpl"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/data"
	"log"
)

type AuthImpl struct {
	AuthCase usecaseimpl.IAuthUseCase
}

func (a *AuthImpl) New() {
	var repository infraAuthRepository.SqlxAuthRepository

	repository.New()

	db, err := data.GetConnection()
	if err != nil {
		log.Println("error to connect db  --> ", err)
	}

	a.AuthCase = &usecaseimpl.Auth{
		Repository: repository,
		Db:         db,
	}

}
