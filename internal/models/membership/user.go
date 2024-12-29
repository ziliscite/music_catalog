package membership

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // gorm model handles timestamps
	Email      string `gorm:"unique;not null"`
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	AccessToken string `json:"token"`
}

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
