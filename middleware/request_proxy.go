package middleware

import (
	auth "golang-gateway-microservice-template/authentication"
	"golang-gateway-microservice-template/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var client = &http.Client{}

// Proxy middleware takes the HTTP request and forwards it to the given microservice endpoint.
// Returns the response from the microservices.
func Proxy(serviceURL string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req, err := http.NewRequest(
			ctx.Request().Method,
			serviceURL+ctx.Request().RequestURI,
			ctx.Request().Body,
		)
		if err != nil {
			return err
		}

		// extend the given request header
		req.Header = ctx.Request().Header

		// add request ID
		if ctx.Response().Writer != nil {
			requestID := ctx.Response().Header().Get(echo.HeaderXRequestID)
			req.Header.Set(echo.HeaderXRequestID, requestID)
		}

		// add user
		user := ctx.Request().Header.Get(auth.HeaderAuthenticatedUser)
		if user != "" {
			req.Header.Set(auth.HeaderAuthenticatedUser, user)
		}

		resp, err := client.Do(req)
		if err != nil {
			ctx.Logger().Errorf("service %s (path %s) could not be reached: %s", serviceURL, ctx.Request().RequestURI, err.Error())
			return utils.Errorf(utils.ErrorTypeBadGateway, "microservice at %s is unavailable", serviceURL)
		}
		defer utils.Close(resp.Body)

		// forward header to response
		for key, value := range resp.Header {
			ctx.Response().Header().Set(key, strings.Join(value, ","))
		}

		contentType := resp.Header.Get(echo.HeaderContentType)
		if contentType == "" {
			contentType = echo.MIMEApplicationJSON
		}

		return ctx.Stream(resp.StatusCode, contentType, resp.Body)
	}
}
