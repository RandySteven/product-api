package handlers

import (
	"net/http"
	"strconv"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/payload/response"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase interfaces.UserUseCase
}

func NewUserHandler(usecase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

// GetAllUsers handles the route for getting all users.
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		utils.ResponseHandler(c.Writer, http.StatusInternalServerError, response.Response{Errors: []string{err.Error()}})
		return
	}

	utils.ResponseHandler(c.Writer, http.StatusOK, response.Response{Message: "Success get all users", Data: users})
}

// GetUserById handles the route for getting a user by ID.
func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseHandler(c.Writer, http.StatusBadRequest, response.Response{Message: "Bad request, invalid id"})
		return
	}

	user, err := h.usecase.GetUserById(uint(id))
	if err != nil {
		utils.ResponseHandler(c.Writer, http.StatusNotFound, response.Response{Errors: []string{"User not found"}})
		return
	}

	utils.ResponseHandler(c.Writer, http.StatusOK, response.Response{Message: "Success get user", Data: user})
}

// Other user-related handlers can be added here as needed.

var _ interfaces.UserHandler = &UserHandler{}
