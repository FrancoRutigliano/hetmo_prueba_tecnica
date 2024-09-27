package authUseCase

import authDomain "hetmo_prueba_tecnica/internal/Auth/pkg/domain"

type AuthUseCase interface {
	Login(authDomain.AuthLoginRequest) (string, error)
}

type AuthUseCaseImpl struct {
}

func NewAuthUseCase() *AuthUseCaseImpl {
	return &AuthUseCaseImpl{}
}

func (a *AuthUseCaseImpl) Login(payload authDomain.AuthLoginRequest) (string, error) {
	return "Login", nil
}

func (a *AuthUseCaseImpl) Register(payload authDomain.AuthRegisterRequest) {

}
