package authController

import (
	authUseCase "hetmo_prueba_tecnica/internal/Auth/pkg/useCases"
)

type Auth struct {
	handler authUseCase.AuthImpl
}

func (a *Auth) NewAuthController() {
	a.handler.New()
}
