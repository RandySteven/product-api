package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handlers) InitRouter(r *mux.Router) {

	productRouter := r.PathPrefix("/products").Subrouter()
	productRouter.Use(h.AuthMiddleware)
	productRouter.HandleFunc("", h.ProductController.CreateProduct).Methods(http.MethodPost)
	productRouter.HandleFunc("", h.ProductController.GetAllProducts).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", h.ProductController.DeleteProductById).Methods(http.MethodDelete)
	productRouter.HandleFunc("/{id}", h.ProductController.GetProductById).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", h.ProductController.UpdateProductById).Methods(http.MethodPut)

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("", h.UserController.GetAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("/register", h.UserController.RegisterUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", h.UserController.GetUserById).Methods(http.MethodGet)
	userRouter.HandleFunc("/login", h.UserController.LoginUser).Methods(http.MethodPost)
}
