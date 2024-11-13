package repository

import (
	"sip/models"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePaymentTx(tx *gorm.DB, payment *models.Payment) error
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

/*
CreatePaymentTx creates a new payment in the database within a transaction.

@param tx *gorm.DB
@param payment *models.Payment

@return error
*/
func (r *paymentRepository) CreatePaymentTx(tx *gorm.DB, payment *models.Payment) error {
	return tx.Create(payment).Error
}
