package usecaseimpl

import (
	"errors"
	"fmt"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	utilsAuth "hetmo_prueba_tecnica/pkg/auth"
	"log"
)

func (a *Auth) Login(payload authDto.AuthLoginRequest) (string, error) {

	user, err := a.Repository.Impl.GetUser(payload, a.Db)
	if err != nil {
		return "", fmt.Errorf("error on db -> %s", err.Error())
	}

	log.Println(user)
	log.Println("payload --> ", payload)

	if !utilsAuth.ComparePasswords(user.Password, []byte(payload.Password)) {
		return "", errors.New("incorrect password")
	}

	return "usuario logeado con exito", nil
}
