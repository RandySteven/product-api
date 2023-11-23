package repository

import (
	"log"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// DeleteAllProducts implements interfaces.ProductRepository.
func (repo *productRepository) DeleteAllProducts() {
	repo.db.Exec("DELETE FROM products")
}

// DeleteProductById implements repositories.productRepository.
func (repo *productRepository) DeleteProductById(id uint) error {
	return repo.db.Exec("UPDATE products SET deleted_at = NOW() WHERE id = ?", id).Error
}

// Find implements repositories.productRepository.
func (repo *productRepository) Find() ([]models.Product, error) {
	var products []models.Product
	if err := repo.db.Where("deleted_at IS NULL").Find(&products).Error; err != nil {
		log.Println("QUERY ERROR : ", err)
		return nil, err
	}
	return products, nil
}

// GetProductById implements repositories.productRepository.
func (repo *productRepository) GetProductById(id uint) (*models.Product, error) {
	var product models.Product
	if err := repo.db.Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Save implements repositories.productRepository.
func (repo *productRepository) Save(product *models.Product) (*models.Product, error) {
	if err := repo.db.Create(product).Error; err != nil {
		log.Println("QUERY ROW ERROR : ", err.Error())
		return nil, err
	}
	return product, nil
}

// UpdateProductById implements repositories.productRepository.
func (repo *productRepository) UpdateProductById(id uint, updatedProduct *models.Product) (*models.Product, error) {
	if err := repo.db.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProduct).Error; err != nil {
		log.Println("QUERY ERR : ", err)
		return nil, err
	}
	return updatedProduct, nil
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

var _ interfaces.ProductRepository = &productRepository{}
