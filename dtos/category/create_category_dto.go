package dtos

import (
	"sip/dtos"
)

type CreateCategoryDTO struct {
	dtos.BaseDTO
	Name string `json:"name" validate:"required,gte=5,lte=50"`
}
