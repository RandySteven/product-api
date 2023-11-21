package repository

import (
	"database/sql"
	"log"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type ProductRepository struct {
	db *sql.DB
}

// DeleteProductById implements repositories.ProductRepository.
func (repo *ProductRepository) DeleteProductById(id uint) error {
	query := "UPDATE products SET deleted_at = NOW() WHERE id = $1"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// Find implements repositories.ProductRepository.
func (repo *ProductRepository) Find() ([]models.Product, error) {
	query := "SELECT id, name, price, stock FROM products WHERE deleted_at IS NULL"
	rows, err := repo.db.Query(query)
	if err != nil {
		log.Println("QUERY ERROR : ", err)
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetProductById implements repositories.ProductRepository.
func (repo *ProductRepository) GetProductById(id uint) (*models.Product, error) {
	query := "SELECT p.id, p.name, p.price, p.stock, c.id, c.name FROM products p JOIN categories c ON p.category_id = c.id WHERE p.id = $1"
	var product models.Product
	product.Category = &models.Category{}
	err := repo.db.QueryRow(query, id).
		Scan(&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Category.ID,
			&product.Category.Name,
		)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Save implements repositories.ProductRepository.
func (repo *ProductRepository) Save(product *models.Product) (*models.Product, error) {
	query := "INSERT INTO products (name, price, stock, category_id, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING ID"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var productId uint
	err = stmt.
		QueryRow(&product.Name, &product.Price, &product.Stock, &product.CategoryID).
		Scan(&productId)
	product.ID = productId
	if err != nil {
		log.Println("QUERY ROW ERROR : ", err.Error())
	}
	return product, nil
}

// UpdateProductById implements repositories.ProductRepository.
func (repo *ProductRepository) UpdateProductById(id uint, product *models.Product) (*models.Product, error) {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4, updated_at=NOW() WHERE id = $5"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Println("QUERY ERR : ", err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&product.Name, &product.Price, &product.Stock, &product.CategoryID, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

var _ interfaces.ProductRepository = &ProductRepository{}
