package handlers_test

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/handlers"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/mocks"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"github.com/go-playground/assert/v2"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

var products = []models.Product{
	{
		ID:         0,
		Name:       "Sabun",
		Price:      3500,
		CategoryID: 1,
		Stock:      999,
	},
	{
		ID:         1,
		Name:       "Sikat Gigi",
		Price:      4000,
		CategoryID: 1,
		Stock:      999,
	},
}

func TestGetAllProducts(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)
		pu.On("GetAllProducts").Return(products, nil)
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/v1/products", nil)
		ph.GetAllProducts(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 500", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)
		pu.On("GetAllProducts").Return(nil, errors.New("test"))
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/v1/products", nil)
		ph.GetAllProducts(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestGetProductById(t *testing.T) {
	t.Run("should return 200", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)

		pu.On("GetProductById", uint(0)).Return(&products[0], nil)
		req, _ := http.NewRequest(http.MethodPost, "/v1/products/1", nil)
		vars := map[string]string{"id": "0"}
		req = mux.SetURLVars(req, vars)

		rec := httptest.NewRecorder()

		ph.GetProductById(rec, req)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 400 due the param variable is invalid", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)

		req, _ := http.NewRequest(http.MethodPost, "/v1/products/invalid_id", nil)
		vars := map[string]string{"id": "invalid_id"}
		req = mux.SetURLVars(req, vars)
		rec := httptest.NewRecorder()

		ph.GetProductById(rec, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 404 due the param variable is invalid", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)

		pu.On("GetProductById", uint(4)).Return(nil, sql.ErrNoRows)
		req, _ := http.NewRequest(http.MethodPost, "/v1/products/invalid_id", nil)
		vars := map[string]string{"id": "4"}
		req = mux.SetURLVars(req, vars)
		rec := httptest.NewRecorder()

		ph.GetProductById(rec, req)
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("should return 500 other error by db or sql", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)

		pu.On("GetProductById", uint(4)).Return(nil, sql.ErrTxDone)
		req, _ := http.NewRequest(http.MethodPost, "/v1/products/invalid_id", nil)
		vars := map[string]string{"id": "4"}
		req = mux.SetURLVars(req, vars)
		rec := httptest.NewRecorder()

		ph.GetProductById(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}

func TestCreateProduct(t *testing.T) {
	t.Run("should return 201", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)
		requestJson := `{
			"name": "Sabun",
			"price": 3500,
			"stock": 999,
			"category_id": 1
		}`

		pu.On("CreateProduct", &products[0]).Return(&products[0], nil)
		req, _ := http.NewRequest(http.MethodPost, "/v1/products", strings.NewReader(requestJson))
		rec := httptest.NewRecorder()

		ph.CreateProduct(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("should return 400", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)
		requestJson := `{
			"price": 3500,
			"stock": 999,
			"category_id": 1
		}`

		product := &models.Product{
			Price:      3500,
			Stock:      999,
			CategoryID: 1,
		}

		pu.On("CreateProduct", product).Return(nil, mock.AnythingOfType("error"))
		req, _ := http.NewRequest(http.MethodPost, "/v1/products", strings.NewReader(requestJson))
		rec := httptest.NewRecorder()

		ph.CreateProduct(rec, req)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 500", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)
		requestJson := `{
			"name": "Sabun",
			"price": 3500,
			"stock": 999,
			"category_id": 1
		}`

		pu.On("CreateProduct", &products[0]).Return(nil, errors.New("Internal server error"))
		req, _ := http.NewRequest(http.MethodPost, "/v1/products", strings.NewReader(requestJson))
		rec := httptest.NewRecorder()

		ph.CreateProduct(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("should return 200 after update product by id", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)
		requestJson := `{
			"name": "Sabun",
			"price": 3500,
			"stock": 999,
			"category_id": 1
		}`

		pu.On("UpdateProductById", uint(0), &models.Product{
			Name:       "Sabun",
			Price:      3500,
			Stock:      999,
			CategoryID: 1,
		}).
			Return(&models.Product{
				Name:       "Sabun",
				Price:      3500,
				Stock:      999,
				CategoryID: 1,
			}, nil)
		req, _ := http.NewRequest(http.MethodPut, "/v1/products/1", strings.NewReader(requestJson))
		rec := httptest.NewRecorder()
		vars := map[string]string{"id": "0"}
		req = mux.SetURLVars(req, vars)

		ph.UpdateProductById(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 400 after update product by id", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)
		requestJson := `{
			"name": "Sabun",
			"price": 3500,
			"stock": 999,
			"category_id": 1
		}`

		req, _ := http.NewRequest(http.MethodPut, "/v1/products/1", strings.NewReader(requestJson))
		rec := httptest.NewRecorder()
		vars := map[string]string{"id": "invalid_id"}
		req = mux.SetURLVars(req, vars)

		ph.UpdateProductById(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 500 after update product by id", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)
		requestJson := `{
			"name": "Sabun",
			"price": 3500,
			"stock": 999,
			"category_id": 1
		}`

		pu.On("UpdateProductById", uint(0), &models.Product{
			Name:       "Sabun",
			Price:      3500,
			Stock:      999,
			CategoryID: 1,
		}).
			Return(nil, sql.ErrNoRows)
		req, _ := http.NewRequest(http.MethodPut, "/v1/products/1", strings.NewReader(requestJson))
		rec := httptest.NewRecorder()
		vars := map[string]string{"id": "0"}
		req = mux.SetURLVars(req, vars)

		ph.UpdateProductById(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("should return 200 after delete product by id", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)

		pu.On("DeleteProductById", uint(1)).Return(nil)
		req, _ := http.NewRequest(http.MethodDelete, "/v1/products/1", nil)
		rec := httptest.NewRecorder()

		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		ph.DeleteProductById(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return 400 delete product invaldi id", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)

		req, _ := http.NewRequest(http.MethodDelete, "/v1/products/1", nil)
		rec := httptest.NewRecorder()

		vars := map[string]string{"id": "invalid_id"}
		req = mux.SetURLVars(req, vars)

		ph.DeleteProductById(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return 404 due product id not found to delete", func(t *testing.T) {
		pu := mocks.NewProductUseCase(t)
		ph := handlers.NewProductHandler(pu)

		pu.On("DeleteProductById", uint(1)).Return(sql.ErrNoRows)
		req, _ := http.NewRequest(http.MethodDelete, "/v1/products/1", nil)
		rec := httptest.NewRecorder()

		vars := map[string]string{"id": "1"}
		req = mux.SetURLVars(req, vars)

		ph.DeleteProductById(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}
