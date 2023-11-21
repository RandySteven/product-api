package models

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   uint   `json:"role_id" validate:"required"`
	Role     Role   `json:"role"`
}
