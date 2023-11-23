package interfaces

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"github.com/gin-gonic/gin"
)

type (
	UserRepository interface {
		Save(user *models.User) (*models.User, error)
		Find() ([]models.User, error)
		GetUserById(id uint) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
	}

	UserUseCase interface {
		CreateUser(user *models.User) (*models.User, error)
		GetAllUsers() ([]models.User, error)
		GetUserById(id uint) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
	}

	UserHandler interface {
		GetAllUsers(c *gin.Context)
		GetUserById(c *gin.Context)
	}
)
