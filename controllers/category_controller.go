package controllers

import (
	"net/http"
	dtos "sip/dtos/category"
	"sip/services"
	util "sip/utils"

	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(service services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: service}
}

/*
API: GET /api/categories
DESC: Get list of categories
*/
func (c *CategoryController) GetCategories(ctx echo.Context) error {
	categoryDTO := new(dtos.ListCategoryDTO)
	categoryDTO.Page, _ = strconv.Atoi(ctx.QueryParam("Page"))
	categoryDTO.PageSize, _ = strconv.Atoi(ctx.QueryParam("PageSize"))

	if err := ctx.Bind(&categoryDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(categoryDTO); err != nil {
		validationErrors := util.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: util.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	category, total, err := c.categoryService.ListCategory(categoryDTO.Page, categoryDTO.PageSize)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, util.CustomResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, util.CustomResponseWithPagination{
		Status:   http.StatusOK,
		Data:     category,
		Page:     categoryDTO.Page,
		PageSize: categoryDTO.PageSize,
		Total:    total,
	})
}

/*
API: POST /api/categories
DESC: Create new category
*/
func (c *CategoryController) Create(ctx echo.Context) error {
	categoryDTO := new(dtos.CreateCategoryDTO)
	err := ctx.Bind(&categoryDTO)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(categoryDTO); err != nil {
		validationErrors := util.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: util.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	category, err := c.categoryService.CreateCategory(categoryDTO.Name)
	if err != nil {
		return ctx.JSON(err.(*util.CustomError).StatusCode, util.CustomResponse{
			Status:  err.(*util.CustomError).StatusCode,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, util.CustomResponse{Status: http.StatusOK, Message: util.MESSAGE_SUCCESS, Data: category})
}

/*
API: PUT /api/categories
DESC: Update category
*/
func (c *CategoryController) Update(ctx echo.Context) error {
	categoryDTO := new(dtos.UpdateCategoryDTO)

	if err := ctx.Bind(&categoryDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(categoryDTO); err != nil {
		validationErrors := util.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: util.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	category, err := c.categoryService.UpdateCategory(categoryDTO)
	if err != nil {
		return ctx.JSON(err.(*util.CustomError).StatusCode, util.CustomResponse{
			Status:  err.(*util.CustomError).StatusCode,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, util.CustomResponse{Status: http.StatusOK, Message: util.MESSAGE_SUCCESS, Data: category})
}

/*
API: DELETE /api/categories
DESC: Delete category
*/
func (c *CategoryController) Delete(ctx echo.Context) error {
	categoryDTO := new(dtos.DeleteCategoryDTO)

	if err := ctx.Bind(&categoryDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(categoryDTO); err != nil {
		validationErrors := util.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, util.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: util.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	err := c.categoryService.DeleteCategory(categoryDTO.ID)
	if err != nil {
		return ctx.JSON(err.(*util.CustomError).StatusCode, util.CustomResponse{
			Status:  err.(*util.CustomError).StatusCode,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(http.StatusOK, util.CustomResponse{Status: http.StatusOK, Message: util.MESSAGE_SUCCESS, Data: nil})
}
