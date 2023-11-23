package usecases

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type authUseCase struct {
	repo interfaces.AuthRepository
}

// LoginUserByEmail implements interfaces.authUseCase.
func (service *authUseCase) LoginUserByEmail(email string) (*models.User, error) {
	return service.repo.LoginUserByEmail(email)
}

// RegisterUser implements interfaces.authUseCase.
func (service *authUseCase) RegisterUser(user *models.User) (*models.User, error) {
	return service.repo.RegisterUser(user)
}

func NewAuthUseCase(repo interfaces.AuthUseCase) *authUseCase {
	return &authUseCase{repo: repo}
}

var _ interfaces.AuthUseCase = &authUseCase{}
