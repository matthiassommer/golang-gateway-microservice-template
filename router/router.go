package router

import (
	"net/http"

	. "golang-gateway-microservice-template/utils"

	gw_middleware "golang-gateway-microservice-template/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	// Router is used as a proxy to map external requests to the internal services
	// that actually handle the requests.
	Router struct {
		client HttpClient
		echo   Echo
	}

	// HttpClient interface for mocking in tests
	HttpClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	// Echo interface for mocking in tests
	Echo interface {
		Start(address string) error
		Group(prefix string, m ...echo.MiddlewareFunc) (g *echo.Group)
		GET(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route
		POST(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route
		Static(prefix, root string) *echo.Route
	}
)

// NewRouter initializes echo and a HTTP client.
// Echo is set up with gzip, cors, a custom error handler and a custom logger.
// A middleware automatically removes trailing slashes, e.g. /project/ becomes /project.
//
// All required routes throughout the gateway backend will be set here, including protection for secured routes.
// The protection is handled by middlewares.
func NewRouter() *Router {
	echo := echo.New()

	if Environment() == ENV_DEV || Environment() == ENV_LOCAL {
		echo.Debug = true
	}

	router := &Router{
		echo: echo,
	}

	echo.HideBanner = true

	echo.Pre(middleware.RemoveTrailingSlash())
	echo.Pre(gw_middleware.RemoveAuthorizedUser)

	echo.Use(ContextMiddleware())
	echo.Use(gw_middleware.CustomLogger())
	echo.Use(middleware.Recover())
	echo.Use(middleware.RequestID())
	echo.Use(middleware.BodyLimit("10MB"))
	echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 4}))
	echo.Use(middleware.CORS())

	echo.HTTPErrorHandler = HTTPErrorHandler
	echo.Logger.SetLevel(0)

	// redirectHTTP()

	router.registerRoutes()

	return router
}

// redirectHTTP listens on port 80 to redirect HTTP to HTTPS.
func redirectHTTP() {
	e := echo.New()
	e.Use(middleware.HTTPSRedirect())
	go func() {
		e.Logger.Fatal(e.Start(":80"))
	}()
}

// Start sets echo to listen on a particular address.
func (router *Router) Start(address string) error {
	return router.echo.Start(address)
}

// Index is the index endpoint handler.
func (router *Router) Index(c echo.Context) error {
	return c.String(http.StatusOK, "gateway service running")
}

// Health returns OK to indicate a running service.
func (router *Router) Health(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
