package models

import "gorm.io/gorm"

// Initialize Category model
type Payment struct {
	gorm.Model
	OrderId int
	Amount  float64
	Status  string
}
