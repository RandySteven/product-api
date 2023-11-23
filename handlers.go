package main

import (
	"encoding/json"
	"net/http"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/configs"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/handlers"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/response"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/usecase"
	"github.com/golang-jwt/jwt/v5"
)

type (
	Handlers struct {
		ProductHandler interfaces.ProductHandler
		UserHandler    interfaces.UserHandler
		AuthHandler    interfaces.AuthHandler
	}
)

func NewHandlers(repo configs.Repository) (*Handlers, error) {
	productUseCase := usecase.NewProductUseCase(repo.ProductRepository)
	userUseCase := usecase.NewUserUseCase(repo.UserRepository)
	authUseCase := usecase.NewAuthUseCase(repo.AuthRepository)

	return &Handlers{
		ProductHandler: handlers.NewProductHandler(productUseCase),
		UserHandler:    handlers.NewUserHandler(userUseCase),
		AuthHandler:    handlers.NewAuthHandler(authUseCase),
	}, nil
}

func (h Handlers) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		cookie, err := req.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				resp := response.Response{
					Errors: []string{"Unauthorized"},
				}
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(resp)
			}
			return
		}
		tokenString := cookie.Value
		claims := &configs.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return configs.JWT_KEY, nil
		})

		if err != nil {
			resp := response.Response{
				Errors: []string{"Unauthorized"},
			}
			json.NewEncoder(res).Encode(resp)
			return
		}

		if !token.Valid {
			resp := response.Response{
				Errors: []string{"Unauthorized"},
			}
			json.NewEncoder(res).Encode(resp)
			return
		}

		next.ServeHTTP(res, req)
	})
}
