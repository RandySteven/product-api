package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/request"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/utils"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase interfaces.ProductUseCase
}

// GetProductById handles the retrieval of a product by ID.
func (handler *ProductHandler) GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, invalid id"})
		return
	}

	product, err := handler.usecase.GetProductById(uint(id))
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success get product", "data": product})
}

// UpdateProductById handles the update of a product by ID.
func (handler *ProductHandler) UpdateProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, invalid id"})
		return
	}

	var productRequest models.Product
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, invalid JSON"})
		return
	}

	productResp, err := handler.usecase.UpdateProductById(uint(id), &productRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success updated product", "data": productResp})
}

// DeleteProductById handles the deletion of a product by ID.
func (handler *ProductHandler) DeleteProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, invalid id"})
		return
	}

	if err := handler.usecase.DeleteProductById(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Product id for %d not found", id)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success deleted"})
}

// CreateProduct handles the creation of a new product.
func (handler *ProductHandler) CreateProduct(c *gin.Context) {
	var request request.ProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, invalid JSON"})
		return
	}

	validationErr := utils.Validate(request)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}

	product := models.Product{
		Name:       request.Name,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: request.CategoryID,
	}

	storeProduct, err := handler.usecase.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Success add product", "data": storeProduct})
}

// GetAllProducts handles the retrieval of all products.
func (handler *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := handler.usecase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success get all products", "data": products})
}

func NewProductHandler(usecase interfaces.ProductUseCase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

var _ interfaces.ProductHandler = &ProductHandler{}
