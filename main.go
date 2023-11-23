package main

import (
	"log"
	"net/http"
	"os"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/configs"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
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
	config := models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)

	service, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer service.Close()

	handlers, err := NewHandlers(*service)
	if err != nil {
		return
	}

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/login", handlers.AuthHandler.LoginUser)
	v1.HandleFunc("/register", handlers.AuthHandler.RegisterUser)
	v1.HandleFunc("/logout", handlers.AuthHandler.LogoutUser)
	handlers.InitRouter(v1)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
	os.Exit(0)
}
