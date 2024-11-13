package controllers

import (
	"net/http"
	dtos "sip/dtos/auth"
	"sip/services"
	"sip/utils"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

/*
API: POST /api/auth/login
DESC: Login user
*/
func (c *AuthController) Login(ctx echo.Context) error {
	userDTO := new(dtos.LoginUserDTO)

	err := ctx.Bind(&userDTO)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: nil})
	}

	//validate
	if err := ctx.Validate(userDTO); err != nil {
		validationErrors := utils.HandleValidationError(err)
		return ctx.JSON(http.StatusBadRequest, utils.CustomResponse{
			Status:  http.StatusBadRequest,
			Message: utils.MESSAGE_VALIDATION_ERROR,
			Data:    validationErrors,
		})
	}

	user, token, err := c.authService.LoginAuthUser(userDTO)
	if err != nil {
		return ctx.JSON(err.(*utils.CustomError).StatusCode, utils.CustomResponse{Status: err.(*utils.CustomError).StatusCode, Message: err.Error(), Data: nil})
	}

	return ctx.JSON(http.StatusOK, utils.CustomResponse{Status: http.StatusOK, Message: utils.MESSAGE_SUCCESS, Data: utils.H{"User": user, "AccessToken": token}})
}
