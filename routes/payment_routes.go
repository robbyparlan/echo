package routes

import (
	"sip/controllers"
	// "sip/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterPaymentRoutes(e *echo.Group, paymentController *controllers.PaymentController) {
	r := e.Group("/payment")
	// r.Use(middlewares.JwtMiddleware())
	r.POST("/paid", paymentController.HandlePaid)
}
