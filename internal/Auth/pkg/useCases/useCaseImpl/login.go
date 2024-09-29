package usecaseimpl

import (
	"errors"
	"fmt"
	authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"
	utilsAuth "hetmo_prueba_tecnica/pkg/auth"
	authJwt "hetmo_prueba_tecnica/pkg/jwt"
	"log"
	"os"
)

func (a *Auth) Login(payload authDto.AuthLoginRequest) (string, error) {

	user, err := a.Repository.Impl.GetUser(payload, a.Db)
	if err != nil {
		return "", fmt.Errorf("error on db -> %s", err.Error())
	}

	log.Println(user)

	if !utilsAuth.ComparePasswords(user.Password, []byte(payload.Password)) {
		return "", errors.New("incorrect password")
	}

	secret := []byte(os.Getenv("SECRET_JWT"))

	token, err := authJwt.Create(secret, user.Id)
	if err != nil {
		return "", fmt.Errorf("error generating jwt --> %s", err.Error())
	}

	return token, nil
}
