package routes

import (
	"sip/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Group, authController *controllers.AuthController) {
	r := e.Group("/auth")
	r.POST("/login", authController.Login)
}
