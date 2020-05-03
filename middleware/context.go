package middleware

import (
	"github.com/labstack/echo/v4"
)

// Context extends echo.Context.
type Context struct {
	echo.Context
}

// ContextMiddleware shows how to extend the echo.Context to Context.
func ContextMiddleware() echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			customContext := &Context{c}
			return handler(customContext)
		}
	}
}
