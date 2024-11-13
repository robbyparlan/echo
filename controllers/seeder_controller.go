package controllers

import (
	"net/http"
	"sip/seeders"
	util "sip/utils"

	"github.com/labstack/echo/v4"
)

type SeederController struct {
}

/*
API: GET /api/seeder
DESC: Seeder
*/
func (c *SeederController) Seeder(ctx echo.Context) error {
	seeders.Seed(&util.DB)

	return ctx.JSON(http.StatusOK, util.CustomResponse{Status: http.StatusOK, Message: util.MESSAGE_SUCCESS, Data: nil})
}
