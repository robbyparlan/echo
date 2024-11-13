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

	e.Use(middleware.RequestID())

	// config cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		MaxAge:       86400, // 1 day
	}))

	e.Use(middleware.CSRF())

	/*
		Secure middleware provides protection against
		cross-site scripting (XSS) attack, content type sniffing, clickjacking,
		insecure connection and other code injection attacks.
		Set By Default Echo
	*/
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		HSTSMaxAge:            3600,
		HSTSExcludeSubdomains: true,
	}))

	// config rate limit
	configRateLimiter := util.GetRateLimiter()
	e.Use(middleware.RateLimiterWithConfig(configRateLimiter))

	routes.RegisterRoutes(e)

	server := &http.Server{
		Addr:         ":" + util.APP_PORT,
		ReadTimeout:  time.Duration(util.APP_READ_TIME_OUT) * time.Second,
		WriteTimeout: time.Duration(util.APP_WRITE_TIME_OUT) * time.Second,
	}
	e.Logger.Fatal(e.StartServer(server))
}
