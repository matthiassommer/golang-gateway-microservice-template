package utils

// AuthorizedUser represents the user shared by the services in the Authorized-User header.
type AuthorizedUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
