package middleware

import (
	"encoding/json"
	auth "golang-gateway-microservice-template/authentication"
	"golang-gateway-microservice-template/utils"

	"github.com/labstack/echo/v4"
)

// Context extends echo.Context.
type Context struct {
	echo.Context
}

// User extracts the user object from the request header.
func (c *Context) User() (*auth.AuthenticatedUser, error) {
	headerUser := c.Request().Header.Get(auth.HeaderAuthenticatedUser)
	if headerUser == "" {
		return nil, utils.Error("send '"+auth.HeaderAuthenticatedUser+"' in the request header", utils.ErrorTypeUnauthorized)
	}

	user := &auth.AuthenticatedUser{}
	err := json.Unmarshal([]byte(headerUser), user)
	if err != nil {
		return nil, utils.Error(err, utils.ErrorTypeUnauthorized)
	}

	return user, nil
}

// ContextMiddleware representa a Echo middleware to extend the echo.Context to commons.Context.
func ContextMiddleware() echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			customContext := &Context{c}
			return handler(customContext)
		}
	}
}
