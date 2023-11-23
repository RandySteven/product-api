package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" gorm:"not null"`
	Price      uint      `json:"price" gorm:"not null"`
	Stock      uint      `json:"stock" gorm:"not null"`
	CategoryID uint      `json:"category_id,omitempty" gorm:"not null"`
	Category   *Category `json:"category,omitempty" gorm:"foreignKey:CategoryID;references:ID"`
	CreatedAt  time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt  gorm.DeletedAt
}
