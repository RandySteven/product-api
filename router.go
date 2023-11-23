package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handlers) InitRouter(r *mux.Router) {

	productRouter := r.PathPrefix("/products").Subrouter()
	productRouter.Use(h.AuthMiddleware)
	productRouter.HandleFunc("", h.ProductHandler.CreateProduct).Methods(http.MethodPost)
	productRouter.HandleFunc("", h.ProductHandler.GetAllProducts).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", h.ProductHandler.DeleteProductById).Methods(http.MethodDelete)
	productRouter.HandleFunc("/{id}", h.ProductHandler.GetProductById).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", h.ProductHandler.UpdateProductById).Methods(http.MethodPut)

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", h.UserHandler.GetAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", h.UserHandler.GetUserById).Methods(http.MethodGet)
}
