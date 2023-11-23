package main

import (
	"log"
	"net/http"
	"os"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/configs"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env got")
	}
	r := gin.New()

	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	config := models.NewConfig(dbHost, dbPort, dbUser, dbPass, dbName)

	logger := &Logger{}
	logger.ProductLog("logs/user.service.log")
	logger.UserLog("logs/user.service.log")

	service, err := configs.NewRepository(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer service.Close()

	handlers, err := server.NewHandlers(*service)
	if err != nil {
		return
	}

	v2 := r.Group("/v2")
	v2.POST("/login", handlers.AuthHandler.LoginUser)
	v2.POST("/register", handlers.AuthHandler.RegisterUser)
	v2.POST("/logout", handlers.AuthHandler.LogoutUser)
	handlers.InitRouter(v2)

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	srv.ListenAndServe()
	os.Exit(0)
}
