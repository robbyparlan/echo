package container

import (
	"sip/utils"
)

type Container struct {
	CategoryContainer *CategoryContainer
	AuthContainer     *AuthContainer
	PaymentContainer  *PaymentContainer
}

func NewContainer() *Container {
	db := &utils.DB

	categoryContainer := NewCategoryContainer(db)
	authContainer := NewAuthContainer(db)
	paymentContainer := NewPaymentContainer(db)

	return &Container{
		CategoryContainer: categoryContainer,
		AuthContainer:     authContainer,
		PaymentContainer:  paymentContainer,
	}
}
