package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/response"
	"github.com/gorilla/mux"
)

type UserController struct {
	usecase interfaces.UserUseCase
}

// GetAllUsers implements interfaces.UserController.
func (controller *UserController) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	users, err := controller.usecase.GetAllUsers()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		resp := response.Response{
			Errors: []string{err.Error()},
		}
		json.NewEncoder(res).Encode(resp)
		return
	}
	resp := response.Response{
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
		resp := response.Response{
			Message: "Bad request, invalid id",
		}
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(resp)
		return
	}
	user, err := controller.usecase.GetUserById(uint(id))
	if err != nil {
		resp := response.Response{
			Errors: []string{"User not found"},
		}
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(resp)
		return
	}
	resp := response.Response{
		Message: "Success get user",
		Data:    user,
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp)
}

func NewUserHandler(usecase interfaces.UserUseCase) *UserController {
	return &UserController{usecase: usecase}
}

var _ interfaces.UserHandler = &UserController{}
