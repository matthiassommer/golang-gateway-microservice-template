package router

import (
	"golang-gateway-microservice-template/config"
	"golang-gateway-microservice-template/middleware"
)

// RegisterRoutes adds all routes to the router.
func (router *Router) registerRoutes() {
	router.registerGeneralRoutes()
	router.registerPizzaRoutes()
}

func (router *Router) registerGeneralRoutes() {
	// Serve swagger spec
	router.echo.Static("/swagger", "swaggerui")

	router.echo.GET("/", router.Index)
	router.echo.GET("/health", router.Health)
}

func (router *Router) registerPizzaRoutes() {
	v1 := router.echo.Group(config.APIVersion)

	pizza := v1.Group("/pizza")

	// swagger:operation GET /pizza pizza GetPizzas
	// Get a list of all pizzas.
	// ---
	// responses:
	//     200:
	//         description: Complete pizza menu received
	//     500:
	//         '$ref': '#/responses/InternalServer'
	pizza.GET("", middleware.Proxy(config.PizzaServiceURL))

	// swagger:operation GET /pizza/:name pizza GetPizzaByName
	// Get a pizza by Name.
	// ---
	// security:
	// - Bearer: []
	// parameters:
	// - name: name
	//   in: path
	//   description: Pizza name
	//   required: true
	//   type: string
	// responses:
	//     200:
	//         '$ref': '#/responses/PizzaResponse'
	//     400:
	//         '$ref': '#/responses/BadRequest'
	//     401:
	//         '$ref': '#/responses/Unauthorized'
	//     403:
	//         '$ref': '#/responses/Forbidden'
	//     404:
	//         '$ref': '#/responses/NotFound'
	//     500:
	//         '$ref': '#/responses/InternalServer'
	pizza.GET("/:name", middleware.Proxy(config.PizzaServiceURL))

	// swagger:operation POST /pizza pizza AddPizza
	// Add a new pizza to the menu.
	// ---
	// security:
	// - Bearer: []
	// parameters:
	// - in: body
	//   name: pizza
	//   required: true
	//   type: string
	// responses:
	//     201:
	//         '$ref': '#/responses/PizzaResponse'
	//     400:
	//         '$ref': '#/responses/BadRequest'
	//     401:
	//         '$ref': '#/responses/Unauthorized'
	//     403:
	//         '$ref': '#/responses/Forbidden'
	//     500:
	//         '$ref': '#/responses/InternalServer'
	pizza.POST("", middleware.Proxy(config.PizzaServiceURL))

	// swagger:operation PUT /pizza/:name pizza UpdatePizzaByName
	// Update a pizza by name.
	// ---
	// security:
	// - Bearer: []
	// parameters:
	// - name: name
	//   in: path
	//   description: Pizza name
	//   required: true
	//   type: string
	// - in: body
	//   name: pizza
	//   required: true
	//   type: string
	//   schema:
	//       '$ref': '#/definitions/PizzaResponse'
	// responses:
	//     200:
	//         '$ref': '#/responses/PizzaResponse'
	//     400:
	//         '$ref': '#/responses/BadRequest'
	//     401:
	//         '$ref': '#/responses/Unauthorized'
	//     403:
	//         '$ref': '#/responses/Forbidden'
	//     404:
	//         '$ref': '#/responses/NotFound'
	//     500:
	//         '$ref': '#/responses/InternalServer'
	pizza.PATCH("/:name", middleware.Proxy(config.PizzaServiceURL))

	// swagger:operation DELETE /pizza/:name pizza DeletePizzaByName
	// Delete a pizza by name.
	// ---
	// security:
	// - Bearer: []
	// parameters:
	// - name: name
	//   in: path
	//   description: Pizza name
	//   required: true
	//   type: string
	// responses:
	//     200:
	//         description: Pizza removed from menu
	//     400:
	//         '$ref': '#/responses/BadRequest'
	//     401:
	//         '$ref': '#/responses/Unauthorized'
	//     403:
	//         '$ref': '#/responses/Forbidden'
	//     404:
	//         '$ref': '#/responses/NotFound'
	//     500:
	//         '$ref': '#/responses/InternalServer'
	pizza.DELETE("/:name", middleware.Proxy(config.PizzaServiceURL))
}
