package interfaces

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"github.com/gin-gonic/gin"
)

type (
	ProductRepository interface {
		Save(product *models.Product) (*models.Product, error)
		Find() ([]models.Product, error)
		GetProductById(id uint) (*models.Product, error)
		UpdateProductById(id uint, product *models.Product) (*models.Product, error)
		DeleteProductById(id uint) error
		DeleteAllProducts()
	}

	ProductUseCase interface {
		CreateProduct(product *models.Product) (*models.Product, error)
		GetAllProducts() ([]models.Product, error)
		GetProductById(id uint) (*models.Product, error)
		DeleteProductById(id uint) error
		UpdateProductById(id uint, product *models.Product) (*models.Product, error)
	}

	ProductHandler interface {
		CreateProduct(c *gin.Context)
		GetAllProducts(c *gin.Context)
		GetProductById(c *gin.Context)
		DeleteProductById(c *gin.Context)
		UpdateProductById(c *gin.Context)
	}
)
