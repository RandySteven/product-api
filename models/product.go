package models

import "github.com/go-playground/validator/v10"

type Product struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name" validate:"required,min=3,max=32"`
	Price      uint      `json:"price" validate:"required"`
	Stock      uint      `json:"stock" validate:"required"`
	CategoryID uint      `json:"category_id,omitempty" validate:"required"`
	Category   *Category `json:"category,omitempty"`
}

func (p *Product) Validate() []validator.FieldError {
	validate := validator.New()
	err := validate.Struct(p)
	return err.(validator.ValidationErrors)
}

type ProductResponse struct {
	ID       uint
	Name     string
	Price    uint
	Stock    uint
	Category Category
}
