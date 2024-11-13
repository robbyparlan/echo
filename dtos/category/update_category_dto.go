package dtos

import (
	"sip/dtos"
)

type UpdateCategoryDTO struct {
	dtos.BaseDTO
	ID   int    `json:"ID" validate:"required"`
	Name string `json:"Name" validate:"required,gte=5,lte=50"`
}
