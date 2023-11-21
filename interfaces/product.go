package interfaces

import (
	"net/http"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type (
	ProductRepository interface {
		Save(product *models.Product) (*models.Product, error)
		Find() ([]models.Product, error)
		GetProductById(id uint) (*models.Product, error)
		UpdateProductById(id uint, product *models.Product) (*models.Product, error)
		DeleteProductById(id uint) error
	}

	ProductService interface {
		CreateProduct(product *models.Product) (*models.Product, error)
		GetAllProducts() ([]models.Product, error)
		GetProductById(id uint) (*models.Product, error)
		DeleteProductById(id uint) error
		UpdateProductById(id uint, product *models.Product) (*models.Product, error)
	}

	ProductController interface {
		CreateProduct(res http.ResponseWriter, req *http.Request)
		GetAllProducts(res http.ResponseWriter, req *http.Request)
		GetProductById(res http.ResponseWriter, req *http.Request)
		DeleteProductById(res http.ResponseWriter, req *http.Request)
		UpdateProductById(res http.ResponseWriter, req *http.Request)
	}
)
