package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          uint           `json:"id" example:"1"`
	FirstName   string         `json:"first_name" example:"Rohit"`
	LastName    string         `json:"last_name" example:"Laishram"`
	DateOfBirth time.Time      `json:"date_of_birth" example:"2002-04-15T00:00:00Z"`
	Email       string         `json:"email" example:"rohit@example.com"`
	Address     string         `json:"address" example:"Imphal"`
	CreatedAt   time.Time      `json:"created_at" example:"2025-02-24T06:50:58.947537634Z"`
	UpdatedAt   time.Time      `json:"updated_at" example:"2025-02-24T06:50:58.947537634Z"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}
