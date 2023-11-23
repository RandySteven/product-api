package server

import (
	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.RouterGroup) {

	productRouter := r.Group("/products")
	// productRouter.Use(h.AuthMiddleware)
	productRouter.POST("", h.ProductHandler.CreateProduct)
	productRouter.GET("", h.ProductHandler.GetAllProducts)
	productRouter.DELETE("/:id", h.ProductHandler.DeleteProductById)
	productRouter.GET("/:id", h.ProductHandler.GetProductById)
	productRouter.PUT("/:id", h.ProductHandler.UpdateProductById)

	// userRouter := r.PathPrefix("/users").Subrouter()
	// userRouter.Use(h.RoleMiddleware)
	// userRouter.HandleFunc("", h.UserHandler.GetAllUsers).Methods(http.MethodGet)
	// userRouter.HandleFunc("/{id}", h.UserHandler.GetUserById).Methods(http.MethodGet)
}
