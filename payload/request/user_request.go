package request

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,alphaunicode"`
}

type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	RoleID   uint   `json:"role_id" validate:"required"`
}
