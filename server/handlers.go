package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/configs"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/enums"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/handlers"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/response"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/usecases"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	Handlers struct {
		ProductHandler interfaces.ProductHandler
		UserHandler    interfaces.UserHandler
		AuthHandler    interfaces.AuthHandler
	}
)

func NewHandlers(repo configs.Repository) (*Handlers, error) {
	productService := usecases.NewProductUseCase(repo.ProductRepository)
	userService := usecases.NewUserUseCase(repo.UserRepository)
	authService := usecases.NewAuthUseCase(repo.AuthRepository)

	return &Handlers{
		ProductHandler: handlers.NewProductHandler(productService),
		UserHandler:    handlers.NewUserHandler(userService),
		AuthHandler:    handlers.NewAuthHandler(authService),
	}, nil
}

func (h Handlers) RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		xRequestId := uuid.New().String()
		xTimestamp := time.Now()
		rctx := context.WithValue(req.Context(), enums.XRequestID, xRequestId)
		rctx = context.WithValue(rctx, enums.XRequestID, xRequestId)
		rctx2 := context.WithValue(rctx, enums.XTimestamp, xTimestamp)

		log.Printf("[Request] %v %v %s \n", req.Method, req.URL.Path, xRequestId)
		log.Printf("[Time] %v ", xTimestamp)

		next.ServeHTTP(res, req.WithContext(rctx2))
	})
}

func (h Handlers) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Println("Log")
		next.ServeHTTP(res, req)
	})
}

func (h Handlers) RoleMiddleware(c *gin.Context) {
	claims := h.validateToken(c)
	if claims == nil {
		resp := response.Response{
			Errors: []string{"Unauthorized. Invalid token"},
		}
		utils.ResponseHandler(c.Writer, http.StatusUnauthorized, resp)
		c.Abort()
		return
	}

	if claims.User != nil && claims.User.RoleID != 1 {
		resp := response.Response{
			Errors: []string{"Access denied"},
		}
		utils.ResponseHandler(c.Writer, http.StatusForbidden, resp)
		c.Abort()
		return
	}

	c.Next()

}

func (h Handlers) validateToken(c *gin.Context) *configs.JWTClaim {
	session := sessions.Default(c)
	tokenString := session.Get("token")
	if tokenString == nil {
		return nil
	}

	claims := &configs.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString.(string), claims, func(t *jwt.Token) (interface{}, error) {
		return configs.JWT_KEY, nil
	})

	if err != nil || !token.Valid {
		return nil
	}

	return claims
}

func (h Handlers) AuthMiddleware(c *gin.Context) {
	claims := h.validateToken(c)
	if claims == nil {
		resp := response.Response{
			Errors: []string{"Unauthorized. Invalid token"},
		}
		utils.ResponseHandler(c.Writer, http.StatusUnauthorized, resp)
		c.Abort()
		return
	}

	c.Set("user", claims.User)
	c.Next()
}
