package authUseCase

type AuthUseCase interface {
	Login() (string, error)
}

type AuthUseCaseImpl struct {
}

func NewAuthUseCase() *AuthUseCaseImpl {
	return &AuthUseCaseImpl{}
}

func (a *AuthUseCaseImpl) Login() (string, error) {
	return "Login", nil
}
