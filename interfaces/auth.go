package interfaces

import (
	"net/http"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type (
	AuthHandler interface {
		RegisterUser(res http.ResponseWriter, req *http.Request)
		LoginUser(res http.ResponseWriter, req *http.Request)
		LogoutUser(res http.ResponseWriter, req *http.Request)
	}

	AuthRepository interface {
		LoginUserByEmail(email string) (*models.User, error)
		RegisterUser(user *models.User) (*models.User, error)
	}

	AuthUseCase interface {
		LoginUserByEmail(email string) (*models.User, error)
		RegisterUser(user *models.User) (*models.User, error)
	}
)
