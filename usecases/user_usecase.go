package usecases

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type userUseCase struct {
	repo interfaces.UserRepository
}

// CreateUser implements interfaces.userUseCase.
func (service *userUseCase) CreateUser(user *models.User) (*models.User, error) {
	return service.repo.Save(user)
}

// GetAllUsers implements interfaces.userUseCase.
func (service *userUseCase) GetAllUsers() ([]models.User, error) {
	return service.repo.Find()
}

// GetUserByEmailAndPassword implements interfaces.userUseCase.
func (service *userUseCase) GetUserByEmail(email string) (*models.User, error) {
	return service.repo.GetUserByEmail(email)
}

// GetUserById implements interfaces.userUseCase.
func (service *userUseCase) GetUserById(id uint) (*models.User, error) {
	return service.repo.GetUserById(id)
}

func NewUserUseCase(repo interfaces.UserRepository) *userUseCase {
	return &userUseCase{repo: repo}
}

var _ interfaces.UserUseCase = &userUseCase{}
