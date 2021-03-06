package middleware

import (
	"golang-gateway-microservice-template/authentication"

	"github.com/labstack/echo/v4"
)

// RemoveAuthorizedUser middleware removes the custom header "Authorized-User" to
// prevent clients from passing this header to the gateway.
func RemoveAuthorizedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Del(authentication.HeaderAuthenticatedUser)
		return next(c)
	}
}
