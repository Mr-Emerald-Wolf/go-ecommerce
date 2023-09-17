package models

import (
	"time"
)

// User represents a user in the application.
type User struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"username" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	IsAdmin   bool   `json:"isAdmin" gorm:"not null"`
	// Add more fields as needed, such as Name, Password, Role, etc.
}

type CreateUserSchema struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password,omitempty"`
}
