package models

import "gorm.io/gorm"

// Initialize Users model
type Users struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `json:"Password,omitempty"`
}
