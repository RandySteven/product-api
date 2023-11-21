package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handlers) InitRouter(r *mux.Router) {

	productRouter := r.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", h.ProductController.CreateProduct).Methods(http.MethodPost)
	productRouter.HandleFunc("", h.ProductController.GetAllProducts).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", h.ProductController.DeleteProductById).Methods(http.MethodDelete)
	productRouter.HandleFunc("/{id}", h.ProductController.GetProductById).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", h.ProductController.UpdateProductById).Methods(http.MethodPut)

}
