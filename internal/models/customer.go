package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type Customer struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` // Soft delete field
}

func (Customer) TableName() string {
	return "customers" // Specify your table name here
}
