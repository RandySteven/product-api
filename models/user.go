package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" validate:"required" gorm:"not null"`
	Email     string    `json:"email" validate:"required" gorm:"not null"`
	Password  string    `json:"password" validate:"required" gorm:"not null"`
	RoleID    uint      `json:"role_id" validate:"required" gorm:"not null"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt time.Time `gorm:"not null;default:current_timestamp"`
	DeletedAt gorm.DeletedAt
}
