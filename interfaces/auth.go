package interfaces

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"github.com/gin-gonic/gin"
)

type (
	AuthHandler interface {
		RegisterUser(c *gin.Context)
		LoginUser(c *gin.Context)
		LogoutUser(c *gin.Context)
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
