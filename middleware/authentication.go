package middleware

import (
	"encoding/json"
	auth "golang-gateway-microservice-template/authentication"
	"golang-gateway-microservice-template/utils"

	"github.com/labstack/echo/v4"
)

var dummyUser = &auth.AuthenticatedUser{ID: 1, Username: "gopher", Email: "golang@microservices.com"}

// Authenticate set a dummy user in the request header.
// Replace this with an actual call to an Authentication Provider.
func Authenticate(next echo.HandlerFunc) echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			user, err := json.Marshal(dummyUser)
			if err != nil {
				utils.Error("cannot marshal user object", utils.ErrorTypeInternalServer)
			}

			ctx.Request().Header.Set(auth.HeaderAuthenticatedUser, string(user))

			return next(ctx)
		}
	}
}
