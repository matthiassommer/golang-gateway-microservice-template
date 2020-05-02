package utils

import (
	"encoding/json"

	"github.com/labstack/echo/v4"
)

// HeaderAuthorizedUser is the header parameter used to exchange the user object between microservices.
const HeaderAuthorizedUser = "Authorized-User"

// Context extended echo.Context.
type Context struct {
	echo.Context
}

// User extracts the user object from the request header.
func (c *Context) User() (*AuthorizedUser, error) {
	auth := c.Request().Header.Get(HeaderAuthorizedUser)
	if auth == "" {
		return nil, Error("Please provide HTTP Header '"+HeaderAuthorizedUser+"'", ErrorTypeUnauthorized)
	}

	user := &AuthorizedUser{}
	err := json.Unmarshal([]byte(auth), user)
	if err != nil {
		return nil, Error(err, ErrorTypeUnauthorized)
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
