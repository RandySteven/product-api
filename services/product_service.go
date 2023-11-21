package services

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type ProductService struct {
	repo interfaces.ProductRepository
}

// GetProductById implements interfaces.ProductService.
func (service *ProductService) GetProductById(id uint) (*models.Product, error) {
	return service.repo.GetProductById(id)
}

// UpdateProductById implements interfaces.ProductService.
func (service *ProductService) UpdateProductById(id uint, product *models.Product) (*models.Product, error) {
	return service.repo.UpdateProductById(id, product)
}

// DeleteProductById implements interfaces.ProductService.
func (service *ProductService) DeleteProductById(id uint) error {
	return service.repo.DeleteProductById(id)
}

// CreateProduct implements services.ProductService.
func (service *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	return service.repo.Save(product)
}

// GetAllProducts implements services.ProductService.
func (service *ProductService) GetAllProducts() ([]models.Product, error) {
	return service.repo.Find()
}

func NewProductService(repo interfaces.ProductRepository) *ProductService {
	return &ProductService{repo}
}

var _ interfaces.ProductService = &ProductService{}
