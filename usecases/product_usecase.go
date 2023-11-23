package usecases

import (
	"database/sql"
	"time"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type productUseCase struct {
	repo interfaces.ProductRepository
}

// GetProductById implements interfaces.productUseCase.
func (service *productUseCase) GetProductById(id uint) (*models.Product, error) {
	product, err := service.repo.GetProductById(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, sql.ErrNoRows
		default:
			return nil, err
		}
	}
	return product, nil
}

// UpdateProductById implements interfaces.productUseCase.
func (service *productUseCase) UpdateProductById(id uint, product *models.Product) (*models.Product, error) {
	return service.repo.UpdateProductById(id, product)
}

// DeleteProductById implements interfaces.productUseCase.
func (service *productUseCase) DeleteProductById(id uint) error {
	return service.repo.DeleteProductById(id)
}

// CreateProduct implements services.productUseCase.
func (service *productUseCase) CreateProduct(product *models.Product) (*models.Product, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	return service.repo.Save(product)
}

// GetAllProducts implements services.productUseCase.
func (service *productUseCase) GetAllProducts() ([]models.Product, error) {
	return service.repo.Find()
}

func NewProductUseCase(repo interfaces.ProductRepository) *productUseCase {
	return &productUseCase{repo}
}

var _ interfaces.ProductUseCase = &productUseCase{}
