package interfaces

import (
	"net/http"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type (
	UserRepository interface {
		Save(user *models.User) (*models.User, error)
		Find() ([]models.User, error)
		GetUserById(id uint) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
	}

	UserService interface {
		CreateUser(user *models.User) (*models.User, error)
		GetAllUsers() ([]models.User, error)
		GetUserById(id uint) (*models.User, error)
		GetUserByEmail(email string) (*models.User, error)
	}

	UserController interface {
		GetAllUsers(res http.ResponseWriter, req *http.Request)
		GetUserById(res http.ResponseWriter, req *http.Request)
	}
)
