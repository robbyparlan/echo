package routes

import (
	"sip/container"
	"sip/utils"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	c := container.NewContainer()
	r := e.Group(utils.API_VERSION)

	/*
		Register routes
	*/
	SeederRoutes(r)

	RegisterCategoryRoutes(r, c.CategoryContainer.CategoryController)
	RegisterAuthRoutes(r, c.AuthContainer.AuthController)
	RegisterPaymentRoutes(r, c.PaymentContainer.PaymentController)
}
