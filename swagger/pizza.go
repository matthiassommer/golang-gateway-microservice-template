package swagger

// Contains schema definitions for pizza endpoints.

// The pizza object as returned by the pizza microservice.
// swagger:response PizzaResponse
type PizzaResponse struct {
	// in:body
	pizza struct {
		Pizza pizza `json: "pizza"`
	}
}

// The pizza has a name and contains ingredients.
type pizza struct {
	Name        string       `json:"name"`
	Ingredients []ingredient `json:"ingredients"`
}

// The ingredient and its count.
type ingredient struct {
	Name  string
	Count int
}
