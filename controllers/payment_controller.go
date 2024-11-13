package controllers

import (
	"net/http"
	dtos "sip/dtos/payment"
	"sip/services"
	util "sip/utils"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(paymentService services.PaymentService) *PaymentController {
	return &PaymentController{paymentService: paymentService}
}

func (c *PaymentController) HandlePaid(ctx echo.Context) error {
	paymentDTO := new(dtos.CreatePaymentDTO)

	if err := ctx.Bind(&paymentDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(paymentDTO); err != nil {
		validationErrors := util.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: util.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	payment, err := c.paymentService.CreatePaymentTx(paymentDTO)
	if err != nil {
		return ctx.JSON(err.(*util.CustomError).StatusCode, util.CustomResponse{
			Status:  err.(*util.CustomError).StatusCode,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, util.CustomResponse{Status: http.StatusOK, Message: util.MESSAGE_SUCCESS, Data: payment})

}
