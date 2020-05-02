// Package swagger contains API model definitions, such as responses and body formats.
// https://swagger.io/docs/specification/2-0/describing-responses/
package swagger

type ErrorBody struct {
	Type      string `json:"type"`
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}

// Bad Request error is returned when the request was malformed.
// swagger:response BadRequest
type errBadRequest struct {
	// in:body
	Body ErrorBody
}

// Unauthorized error is returned when the user has sent invalid credentials.
// swagger:response Unauthorized
type errUnauthorized struct {
	// in:body
	Body ErrorBody
}

// Forbidden error is returned when the user is not authorized to access a resource.
// swagger:response Forbidden
type errForbidden struct {
	// in:body
	Body ErrorBody
}

// Resource not found error is returned when the resource was not found.
// swagger:response NotFound
type errNotFound struct {
	// in:body
	Body ErrorBody
}

// Internal server error is returned when the server encountered an unexpected condition that prevented it from fulfilling the request.
// swagger:response InternalServer
type errInternal struct {
	// in:body
	Body ErrorBody
}
