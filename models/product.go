package models

type Product struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Price      uint      `json:"price"`
	Stock      uint      `json:"stock"`
	CategoryID uint      `json:"category_id,omitempty"`
	Category   *Category `json:"category,omitempty"`
}
