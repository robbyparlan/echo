package main

import (
	"net/http"
	util "sip/utils"
	"time"

	"sip/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = util.HTTPErrorHandler

	routes.RegisterRoutes(e)

	server := &http.Server{
		Addr:         ":" + util.APP_PORT,
		ReadTimeout:  time.Duration(util.APP_READ_TIME_OUT) * time.Second,
		WriteTimeout: time.Duration(util.APP_WRITE_TIME_OUT) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(server))
}
