package services

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type AuthService struct {
	repo interfaces.AuthRepository
}

// LoginUserByEmail implements interfaces.AuthService.
func (service *AuthService) LoginUserByEmail(email string) (*models.User, error) {
	return service.repo.LoginUserByEmail(email)
}

// RegisterUser implements interfaces.AuthService.
func (service *AuthService) RegisterUser(user *models.User) (*models.User, error) {
	return service.repo.RegisterUser(user)
}

func NewAuthService(repo interfaces.AuthService) *AuthService {
	return &AuthService{repo: repo}
}

var _ interfaces.AuthService = &AuthService{}
