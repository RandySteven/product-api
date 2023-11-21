package main

import (
	"log"
	"net/http"
	"os"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/controller"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/infrastructure/persistence"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	r := mux.NewRouter()

	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	service, err := persistence.NewRepository(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		log.Println(err)
		return
	}
	defer service.Close()

	productController := controller.NewProductController(services.NewProductService(service.ProductRepository))

	productRouter := r.PathPrefix("/products").Subrouter()
	productRouter.HandleFunc("", productController.CreateProduct).Methods(http.MethodPost)
	productRouter.HandleFunc("", productController.GetAllProducts).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", productController.DeleteProductById).Methods(http.MethodDelete)
	productRouter.HandleFunc("/{id}", productController.GetProductById).Methods(http.MethodGet)
	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
	os.Exit(0)
}
