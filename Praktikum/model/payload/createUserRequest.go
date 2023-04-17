package payload

import "gorm.io/gorm"

type CreateUserRequest struct {
	gorm.Model
	Name     string `json:"name" validate:"required, min= 3 max=20"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required min=8"`
}
