package models

import "gorm.io/gorm"

// Initialize Category model
type Category struct {
	gorm.Model
	Name string
}
