package main

import (
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/controller"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/infrastructure/persistence"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/services"
)

type (
	Handlers struct {
		ProductController interfaces.ProductController
	}
)

func NewHandlers(repo persistence.Repository) (*Handlers, error) {
	productService := services.NewProductService(repo.ProductRepository)

	return &Handlers{
		ProductController: controller.NewProductController(productService),
	}, nil
}
