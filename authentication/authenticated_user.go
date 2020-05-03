package authentication

// AuthenticatedUser represents the user shared by the microservices in the Authorized-User header.
type AuthenticatedUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
