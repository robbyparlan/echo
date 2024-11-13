package routes

import (
	"sip/controllers"
	"sip/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterCategoryRoutes(e *echo.Group, categoryController *controllers.CategoryController) {
	r := e.Group("/category")
	r.Use(middlewares.JwtMiddleware())
	r.GET("", categoryController.GetCategories)
	r.POST("", categoryController.Create)
	r.PUT("", categoryController.Update)
	r.DELETE("", categoryController.Delete)
}
