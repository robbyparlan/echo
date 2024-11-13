package services

import (
	"context"
	"net/http"
	pb "sip/config/protos"
	dtos "sip/dtos/payment"
	"sip/models"
	"sip/repository"
	"sip/utils"

	"gorm.io/gorm"
)

type PaymentService interface {
	CreatePaymentTx(payment *dtos.CreatePaymentDTO) (*models.Payment, error)
}

type paymentService struct {
	repo repository.PaymentRepository
	db   *gorm.DB // Tambahkan *gorm.DB untuk transaksi
}

func NewPaymentService(repo repository.PaymentRepository, db *gorm.DB) PaymentService {
	return &paymentService{
		repo: repo,
		db:   db,
	}
}

func (s *paymentService) CreatePaymentTx(paymentDto *dtos.CreatePaymentDTO) (*models.Payment, error) {
	payment := models.Payment{
		OrderId: paymentDto.OrderId,
		Amount:  paymentDto.Amount,
		Status:  utils.PAYMENT_STATUS_PAID,
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	err := s.repo.CreatePaymentTx(tx, &payment)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	if utils.GRPCClient == nil {
		return nil, utils.NewCustomError(http.StatusInternalServerError, "gRPC client not initialized", nil)
	}

	_, err = (*utils.GRPCClient).UpdatePaymentStatus(context.Background(), &pb.PaymentStatusRequest{
		OrderId: int32(payment.OrderId),
		Status:  payment.Status,
	})
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
