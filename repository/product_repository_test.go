package repository_test

import (
	"log"
	"os"
	"testing"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/infrastructure/persistence"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestProductRepositoryByTestDB(t *testing.T) {
	t.Run("the delete product success should return no error", func(t *testing.T) {
		if err := godotenv.Load(); err != nil {
			log.Println("no env got")
		}
		testConfig := models.Config{
			DbHost: os.Getenv("TEST_DB_HOST"),
			DbName: os.Getenv("TEST_DB_NAME"),
			DbPort: os.Getenv("TEST_DB_PORT"),
			DbUser: os.Getenv("TEST_DB_USER"),
			DbPass: os.Getenv("TEST_DB_PASS"),
		}
		service, _ := persistence.NewRepository(&testConfig)
		product := &models.Product{
			Name:       "TestProduct1",
			Price:      10000,
			Stock:      100,
			CategoryID: 2,
		}
		service.ProductRepository.Save(product)

		err := service.ProductRepository.DeleteProductById(1)
		assert.Nil(t, err)
		service.ProductRepository.DeleteAllProducts()
	})

	t.Run("should return all products by method find", func(t *testing.T) {
		if err := godotenv.Load(); err != nil {
			log.Println("no env got")
		}
		testConfig := models.Config{
			DbHost: os.Getenv("TEST_DB_HOST"),
			DbName: os.Getenv("TEST_DB_NAME"),
			DbPort: os.Getenv("TEST_DB_PORT"),
			DbUser: os.Getenv("TEST_DB_USER"),
			DbPass: os.Getenv("TEST_DB_PASS"),
		}
		product1 := &models.Product{
			Name:       "TestProduct1",
			Price:      10000,
			Stock:      100,
			CategoryID: 1,
		}
		product2 := &models.Product{
			Name:       "TestProduct2",
			Price:      10000,
			Stock:      100,
			CategoryID: 2,
		}
		service, _ := persistence.NewRepository(&testConfig)
		service.ProductRepository.Save(product1)
		service.ProductRepository.Save(product2)
		products, err := service.ProductRepository.Find()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(products))
		service.ProductRepository.DeleteAllProducts()
	})

	t.Run("should return product after created product", func(t *testing.T) {
		if err := godotenv.Load(); err != nil {
			log.Println("no env got")
		}
		testConfig := models.Config{
			DbHost: os.Getenv("TEST_DB_HOST"),
			DbName: os.Getenv("TEST_DB_NAME"),
			DbPort: os.Getenv("TEST_DB_PORT"),
			DbUser: os.Getenv("TEST_DB_USER"),
			DbPass: os.Getenv("TEST_DB_PASS"),
		}
		product := &models.Product{
			Name:       "TestProduct1",
			Price:      10000,
			Stock:      100,
			CategoryID: 1,
		}

		service, _ := persistence.NewRepository(&testConfig)
		savedProduct, _ := service.ProductRepository.Save(product)

		assert.Equal(t, product.Name, savedProduct.Name)
		assert.Equal(t, product.Price, savedProduct.Price)
		assert.Equal(t, product.Stock, savedProduct.Stock)
		assert.Equal(t, product.CategoryID, savedProduct.CategoryID)
		service.ProductRepository.DeleteAllProducts()
	})
}

func TestProductRepositoryByMock(t *testing.T) {
	t.Run("the delete product by id must called", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		query := "^UPDATE products SET deleted_at = NOW\\(\\) WHERE id = \\$1$"
		if err != nil {
			t.Fatalf("error creating sqlmock: %v", err)
		}
		defer db.Close()

		repo := repository.NewProductRepository(db)

		mock.ExpectPrepare(query).
			ExpectExec().
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err = repo.DeleteProductById(1)

		// Assertions
		assert.NoError(t, err, "expected no error")
		assert.NoError(t, mock.ExpectationsWereMet(), "expected all SQL expectations to be met")
	})

	t.Run("should called save method", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		query := "INSERT INTO products \\(name, price, stock, category_id, created_at, updated_at\\) VALUES (.+) RETURNING ID"
		if err != nil {
			t.Fatalf("error creating sqlmock: %v", err)
		}
		defer db.Close()

		repo := repository.NewProductRepository(db)

		mock.ExpectPrepare(query).
			ExpectQuery().
			WithArgs("TestProduct1", 10000, 100, 2).
			WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow(1))

		product := &models.Product{
			Name:       "TestProduct1",
			Price:      10000,
			Stock:      100,
			CategoryID: 2,
		}

		savedProduct, err := repo.Save(product)

		assert.NoError(t, err, "expected no error")
		assert.NotNil(t, savedProduct, "expected a product to be returned")
		assert.Equal(t, uint(1), savedProduct.ID, "expected the ID to be set")
		assert.NoError(t, mock.ExpectationsWereMet(), "expected all SQL expectations to be met")
	})

	t.Run("should called find method", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		query := "SELECT id, name, price, stock FROM products WHERE deleted_at IS NULL"
		if err != nil {
			t.Fatalf("error creating sqlmock: %v", err)
		}
		defer db.Close()

		repo := repository.NewProductRepository(db)

		mock.ExpectQuery(query).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "stock"}).
				AddRow(1, "Product1", 10.0, 100).
				AddRow(2, "Product2", 20.0, 200))

		products, err := repo.Find()

		assert.NoError(t, err, "expected no error")
		assert.Len(t, products, 2, "expected 2 products")
		assert.Equal(t, uint(1), products[0].ID)
		assert.Equal(t, "Product1", products[0].Name)
		assert.NoError(t, mock.ExpectationsWereMet(), "expected all SQL expectations to be met")
	})

}
