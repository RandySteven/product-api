package repository

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

// RegisterUser implements interfaces.AuthRepository.
func (repo *authRepository) RegisterUser(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// LoginUserByEmail implements interfaces.AuthRepository.
func (repo *authRepository) LoginUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

var _ interfaces.AuthRepository = &authRepository{}
