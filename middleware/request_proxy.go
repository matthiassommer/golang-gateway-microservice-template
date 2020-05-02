package middleware

import (
	. "golang-gateway-microservice-template/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

var client = &http.Client{}

// Proxy middleware takes the HTTP request and forwards it to the given service endpoint.
// Returns the response from the microservices.
func Proxy(serviceURL string) echo.HandlerFunc {
	return func(c echo.Context) error {
		req, err := http.NewRequest(
			c.Request().Method,
			serviceURL+c.Request().RequestURI,
			c.Request().Body,
		)
		if err != nil {
			return err
		}

		req.Header = c.Request().Header

		if c.Response().Writer != nil {
			requestID := c.Response().Header().Get(echo.HeaderXRequestID)
			req.Header.Set(echo.HeaderXRequestID, requestID)
		}

		// add the user to the request header
		user := c.Request().Header.Get(HeaderAuthorizedUser)
		if user != "" {
			req.Header.Set(HeaderAuthorizedUser, user)
		}

		resp, err := client.Do(req)
		if err != nil {
			c.Logger().Errorf("service %s (path %s) could not be reached: %s", serviceURL, c.Request().RequestURI, err.Error())
			return Error("a microservice is unavailable", ErrorTypeBadGateway)
		}
		defer Close(resp.Body)

		// forward all headers to response
		for key, value := range resp.Header {
			c.Response().Header().Set(key, value[0])
		}

		contentType := resp.Header.Get(echo.HeaderContentType)
		if contentType == "" {
			contentType = echo.MIMEApplicationJSON
		}

		return c.Stream(resp.StatusCode, contentType, resp.Body)
	}
}
