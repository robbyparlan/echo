// middlewares/auth.go
package middlewares

import (
	"sip/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// BasicAuthMiddleware provides a reusable basic authentication middleware
func BasicAuthMiddleware() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Replace with your actual username and password validation
		if username == utils.BASIC_AUTH_USERNAME && password == utils.BASIC_AUTH_PASSWORD {
			return true, nil
		}
		return false, nil
	})
}
