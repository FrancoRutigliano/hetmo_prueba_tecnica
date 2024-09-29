package usecaseimpl

import authDto "hetmo_prueba_tecnica/internal/Auth/pkg/domain/dto"

func (a *Auth) Login(payload authDto.AuthLoginRequest) (string, error) {

	return "Login", nil
}
