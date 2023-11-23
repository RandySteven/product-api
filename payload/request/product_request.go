package request

type ProductRequest struct {
	Name       string `json:"name" validate:"required"`
	Price      uint   `json:"price" validate:"required"`
	Stock      uint   `json:"stock" validate:"required"`
	CategoryID uint   `json:"category_id" validate:"required"`
}
