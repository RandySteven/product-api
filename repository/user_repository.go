package repository

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Find implements interfaces.userRepository.
func (repo *userRepository) Find() ([]models.User, error) {
	var users []models.User

	if err := repo.db.Preload("Role").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByEmailAndPassword implements interfaces.userRepository.
func (repo *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserById implements interfaces.userRepository.
func (repo *userRepository) GetUserById(id uint) (*models.User, error) {
	var user models.User

	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Save implements interfaces.userRepository.
func (repo *userRepository) Save(user *models.User) (*models.User, error) {
	if err := repo.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

var _ interfaces.UserRepository = &userRepository{}
