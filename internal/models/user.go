package models

import (
	"time"

	"gorm.io/gorm"
)

// User model represents the users table in the database
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name" `
	Email     string         `gorm:"uniqueIndex;size:191" json:"email" ` // Set to varchar(191)
	Password  string         `json:"password,omitempty"`                 // Password must be at least 6 characters
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName sets the insert table name for this struct type
func (User) TableName() string {
	return "users" // Specify your table name here
}

// User model represents the users table in the database
type UserCreateRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" `
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
}

type UsetrUpdateRequest struct {
	FirstName string `json:"name" binding:"required"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"type"`
}
