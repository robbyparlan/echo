package container

import (
	"sip/controllers"
	"sip/repository"
	"sip/services"

	"gorm.io/gorm"
)

type PaymentContainer struct {
	PaymentController *controllers.PaymentController
}

func NewPaymentContainer(db *gorm.DB) *PaymentContainer {
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(paymentRepo, db)
	paymentController := controllers.NewPaymentController(paymentService)

	return &PaymentContainer{
		PaymentController: paymentController,
	}
}
