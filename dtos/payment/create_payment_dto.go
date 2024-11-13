package dtos

import (
	"sip/dtos"
)

type CreatePaymentDTO struct {
	dtos.BaseDTO
	OrderId int     `json:"OrderId" validate:"required"`
	Amount  float64 `json:"Amount" validate:"required"`
}
