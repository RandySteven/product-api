package models

type Product struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name" validate:"required,min=3,max=32"`
	Price      uint      `json:"price" validate:"required"`
	Stock      uint      `json:"stock" validate:"required"`
	CategoryID uint      `json:"category_id,omitempty" validate:"required"`
	Category   *Category `json:"category,omitempty"`
}

func (p *Product) Validate() {
}

type ProductResponse struct {
	ID       uint
	Name     string
	Price    uint
	Stock    uint
	Category Category
}
