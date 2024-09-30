package userController

import userUseCase "hetmo_prueba_tecnica/internal/User/pkg/useCases"

type User struct {
	handler userUseCase.UserImpl
}

func (u *User) NewUserController() {
	u.handler.New()
}
