package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Role      string    `json:"role" gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}
