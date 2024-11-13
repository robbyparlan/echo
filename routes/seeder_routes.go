package routes

import (
	"sip/controllers"
	"sip/middlewares"

	"github.com/labstack/echo/v4"
)

func SeederRoutes(e *echo.Group) {
	seederController := controllers.SeederController{}
	r := e.Group("/seeder")
	r.Use(middlewares.BasicAuthMiddleware())
	r.GET("", seederController.Seeder)
}
