package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ProductRequest struct {
	Name       string `json:"name" validate:"required"`
	Price      uint   `json:"price" validate:"required"`
	Stock      uint   `json:"stock" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}

func (request *ProductRequest) Validate() []string {
	validate := validator.New()
	err := validate.Struct(request)
	var errs []string
	for _, currErr := range err.(validator.ValidationErrors) {
		errMsg := fmt.Sprintf("%s field is %s", currErr.Field(), currErr.ActualTag())
		errs = append(errs, errMsg)
	}
	return errs
}
