package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/infrastructure/persistence"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/request"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/utils"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	service interfaces.AuthService
}

// LogoutUser implements interfaces.AuthController.
func (*AuthController) LogoutUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	http.SetCookie(res, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})
	resp := models.Response{
		Message: "Success to logout",
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp)
}

func NewAuthController(service interfaces.AuthService) *AuthController {
	return &AuthController{service: service}
}

// LoginUser implements interfaces.AuthController.
func (controller *AuthController) LoginUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var request request.UserLoginRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	pass, err := utils.HashPassword(request.Password)
	request.Password = pass
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	user, err := controller.service.LoginUserByEmail(request.Email)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	// if err != nil {
	// 	res.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	expTime := time.Now().Add(time.Minute * 1)
	claims := &persistence.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(persistence.JWT_KEY)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	http.SetCookie(res, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	resp := models.Response{
		Message: "Success login user",
		Data:    user,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp)
}

// RegisterUser implements interfaces.AuthController.
func (controller *AuthController) RegisterUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}

	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	user.Password = pass
	userStore, err := controller.service.RegisterUser(&user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	resp := models.Response{
		Message: "Success created user",
		Data:    userStore,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp)
}

var _ interfaces.AuthController = &AuthController{}
