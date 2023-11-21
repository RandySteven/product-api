package services

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type UserService struct {
	repo interfaces.UserRepository
}

// CreateUser implements interfaces.UserService.
func (service *UserService) CreateUser(user *models.User) (*models.User, error) {
	return service.repo.Save(user)
}

// GetAllUsers implements interfaces.UserService.
func (service *UserService) GetAllUsers() ([]models.User, error) {
	return service.repo.Find()
}

// GetUserByEmailAndPassword implements interfaces.UserService.
func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	return service.repo.GetUserByEmail(email)
}

// GetUserById implements interfaces.UserService.
func (service *UserService) GetUserById(id uint) (*models.User, error) {
	return service.repo.GetUserById(id)
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

var _ interfaces.UserService = &UserService{}
