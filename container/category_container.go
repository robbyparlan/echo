package container

import (
	"sip/controllers"
	"sip/repository"
	"sip/services"

	"gorm.io/gorm"
)

type CategoryContainer struct {
	CategoryController *controllers.CategoryController
}

func NewCategoryContainer(db *gorm.DB) *CategoryContainer {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo, db)
	categoryController := controllers.NewCategoryController(categoryService)

	return &CategoryContainer{
		CategoryController: categoryController,
	}
}
