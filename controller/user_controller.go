package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/utils"
	"github.com/gorilla/mux"
)

type UserController struct {
	service interfaces.UserService
}

// GetAllUsers implements interfaces.UserController.
func (controller *UserController) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	users, err := controller.service.GetAllUsers()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	resp := models.Response{
		Message: "Success get all users",
		Data:    users,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp)
}

// GetUserById implements interfaces.UserController.
func (controller *UserController) GetUserById(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		resp := models.Response{
			Message: "Bad request, invalid id",
		}
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(resp)
		return
	}
	user, err := controller.service.GetUserById(uint(id))
	if err != nil {
		resp := models.Response{
			Errors: []string{"User not found"},
		}
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(resp)
		return
	}
	resp := models.Response{
		Message: "Success get user",
		Data:    user,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp)
}

// LoginUser implements interfaces.UserController.
func (controller *UserController) LoginUser(res http.ResponseWriter, req *http.Request) {
	panic("unimplemented")
}

// RegisterUser implements interfaces.UserController.
func (controller *UserController) RegisterUser(res http.ResponseWriter, req *http.Request) {
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
	userStore, err := controller.service.CreateUser(&user)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := models.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	json.NewEncoder(res).Encode(userStore)
}

func NewUserController(service interfaces.UserService) *UserController {
	return &UserController{service: service}
}

var _ interfaces.UserController = &UserController{}
