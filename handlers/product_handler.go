package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/request"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/response"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/utils"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	usecase interfaces.ProductUseCase
}

// GetProductById implements interfaces.ProductHandler.
func (controller *ProductHandler) GetProductById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		resp := response.Response{
			Message: "Bad request, invalid id",
		}
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(resp)
		return
	}
	product, err := controller.usecase.GetProductById(uint(id))
	if err != nil {
		resp := response.Response{
			Message: "Product not found",
		}
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(resp)
		return
	}
	resp := response.Response{
		Message: "Success deleted",
		Data:    product,
	}
	json.NewEncoder(res).Encode(resp)
}

// UpdateProductById implements interfaces.ProductHandler.
func (controller *ProductHandler) UpdateProductById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		resp := response.Response{
			Message: "Bad request, invalid id",
		}
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(resp)
		return
	}
	var productRequest models.Product
	err = json.NewDecoder(req.Body).Decode(&productRequest)
	if err != nil {
		resp := response.Response{
			Message: "Bad Request",
		}
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(resp)
		return
	}
	productResp, err := controller.usecase.UpdateProductById(uint(id), &productRequest)
	if err != nil {
		resp := response.Response{
			Message: "Internal server error",
		}
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(resp)
		return
	}
	resp := response.Response{
		Message: "Success updated product",
		Data:    productResp,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp)

}

// DeleteProductById implements interfaces.ProductHandler.
func (controller *ProductHandler) DeleteProductById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		resp := response.Response{
			Message: "Bad request, invalid id",
		}
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(resp)
		return
	}
	err = controller.usecase.DeleteProductById(uint(id))
	if err != nil {
		resp := response.Response{
			Message: fmt.Sprintf("product id for %d not found", id),
		}
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(resp)
		return
	}
	resp := response.Response{
		Message: "Success deleted",
	}
	json.NewEncoder(res).Encode(resp)
}

// CreateProduct implements interfaces.ProductHandler.
func (controller *ProductHandler) CreateProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var request request.ProductRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		resp := response.Response{
			Message: "Bad request",
			Errors:  []string{err.Error()},
		}
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(resp)
		return
	}
	validationErr := utils.Validate(request)
	if validationErr != nil {
		resp := response.Response{
			Message: "Bad request",
			Errors:  validationErr,
		}
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(resp)
		return
	}
	product := models.Product{
		Name:       request.Name,
		Price:      request.Price,
		Stock:      request.Stock,
		CategoryID: request.CategoryID,
	}
	storeProduct, err := controller.usecase.CreateProduct(&product)
	if err != nil {
		resp := response.Response{
			Message: "Internal server error",
		}
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(resp)
		return
	}
	res.WriteHeader(http.StatusCreated)
	resp := response.Response{
		Message: "Success add product",
		Data:    storeProduct,
	}
	json.NewEncoder(res).Encode(resp)
}

// GetAllProducts implements interfaces.ProductHandler.
func (controller *ProductHandler) GetAllProducts(res http.ResponseWriter, req *http.Request) {
	products, err := controller.usecase.GetAllProducts()
	res.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(products)
}

func NewProductHandler(usecase interfaces.ProductUseCase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

var _ interfaces.ProductHandler = &ProductHandler{}
