package middlewares

import (
	"net/http"
	"os"
	"sip/utils"

	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

/* JwtMiddleware provides a reusable JWT middleware */
func JwtMiddleware() echo.MiddlewareFunc {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	config := echoJwt.Config{
		SigningKey: secretKey,
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, utils.CustomResponse{Status: http.StatusUnauthorized, Message: "Unauthorized", Data: nil})
		},
	}

	return echoJwt.WithConfig(config)
}
