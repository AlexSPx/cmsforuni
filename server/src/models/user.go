package models

import (
	"time"

	"github.com/google/uuid"
)

// User database model
type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`

	Username string `gorm:"uniqueIndex;type:varchar(255);not null"`
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Photo    string
	Approved bool `gorm:"not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserSignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Username        string `json:"username" binding:"required,min=6"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Username string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Approved  bool      `json:"approved"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
