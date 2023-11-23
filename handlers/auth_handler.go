package handlers

import (
	"net/http"
	"time"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/configs"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/request"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/response"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	usecase interfaces.AuthUseCase
}

func NewAuthHandler(usecase interfaces.AuthUseCase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

// LogoutUser handles the route for logging out a user.
func (h *AuthHandler) LogoutUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	resp := response.Response{
		Message: "Success to logout",
	}
	c.JSON(http.StatusOK, resp)
}

// LoginUser handles the route for logging in a user.
func (h *AuthHandler) LoginUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var request request.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	pass, err := utils.HashPassword(request.Password)
	request.Password = pass
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	user, err := h.usecase.LoginUserByEmail(request.Email)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	expTime := time.Now().Add(time.Minute * 60)
	claims := &configs.JWTClaim{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(configs.JWT_KEY)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	resp := response.Response{
		Message: "Success login user",
		Data:    user,
	}
	c.JSON(http.StatusOK, resp)
}

// RegisterUser handles the route for registering a new user.
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var register request.UserRegisterRequest
	if err := c.ShouldBindJSON(&register); err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	validationErr := utils.Validate(register)
	if validationErr != nil {
		resp := response.Response{
			Errors: validationErr,
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}

	user := &models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: register.Password,
		RoleID:   register.RoleID,
	}

	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}
	user.Password = pass
	userStore, err := h.usecase.RegisterUser(user)
	if err != nil {
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, resp)
		return
	}
	resp := response.Response{
		Message: "Success created user",
		Data:    userStore,
	}
	c.JSON(http.StatusOK, resp)
}

var _ interfaces.AuthHandler = &AuthHandler{}
