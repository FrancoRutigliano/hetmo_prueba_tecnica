package userUseCase

import (
	infraUserRepository "hetmo_prueba_tecnica/internal/User/infrastructure/repository"
	userUseCaseImpl "hetmo_prueba_tecnica/internal/User/pkg/useCases/useCaseImpl"
	"hetmo_prueba_tecnica/internal/shared/infrastructure/data"
	"log"
)

type UserImpl struct {
	UserCase userUseCaseImpl.IUserUseCase
}

func (u *UserImpl) New() {
	var repository infraUserRepository.SqlxUserRepository

	repository.New()

	db, err := data.GetConnection()
	if err != nil {
		log.Println("error to connect db  --> ", err)
	}

	u.UserCase = &userUseCaseImpl.User{
		Repository: repository,
		Db:         db,
	}
}
