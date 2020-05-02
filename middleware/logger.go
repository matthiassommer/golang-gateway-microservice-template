package middleware

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CustomLogger returns a pre-configured Echo logger.
func CustomLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: skipHealthChecks,
	})
}

// skipHealthChecks is used to avoid logging of health checks.
func skipHealthChecks(c echo.Context) bool {
	return strings.HasSuffix(c.Path(), "health")
}
