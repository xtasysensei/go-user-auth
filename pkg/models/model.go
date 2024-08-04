package models

import "time"

// Define your models here: type struct
type User struct {
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
type RegisterUserPayload struct {
	Username        string `json:"username" validate:"required,alphanum"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=3,max=130"`
	ConfirmPassword string `json:"confirmpassword" validate:"required,min=3,max=130"`
}
type LoginUserPayload struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required"`
}

