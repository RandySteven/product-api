package main

import (
	"encoding/json"
	"net/http"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/controller"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/infrastructure/persistence"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/services"
	"github.com/golang-jwt/jwt/v5"
)

type (
	Handlers struct {
		ProductController interfaces.ProductController
		UserController    interfaces.UserController
	}
)

func NewHandlers(repo persistence.Repository) (*Handlers, error) {
	productService := services.NewProductService(repo.ProductRepository)
	userService := services.NewUserService(repo.UserRepository)

	return &Handlers{
		ProductController: controller.NewProductController(productService),
		UserController:    controller.NewUserController(userService),
	}, nil
}

func (h Handlers) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		cookie, err := req.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				resp := models.Response{
					Errors: []string{"Unauthorized"},
				}
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(resp)
			}
			return
		}
		tokenString := cookie.Value
		claims := &persistence.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return persistence.JWT_KEY, nil
		})

		if err != nil {
			resp := models.Response{
				Errors: []string{"Unauthorized"},
			}
			json.NewEncoder(res).Encode(resp)
		}

		if !token.Valid {
			resp := models.Response{
				Errors: []string{"Unauthorized"},
			}
			json.NewEncoder(res).Encode(resp)
		}

		next.ServeHTTP(res, req)
	})
}
