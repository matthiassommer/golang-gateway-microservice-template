package swagger

// Contains schema definitions for auth endpoints.

// Returns the login url in body. Use this URL to generate an access token.
// swagger:response AuthResponse
type authResponse struct {
	// in:body
	Location string `json:"location"`
}

// Token pair generated using the passed refresh token.
// swagger:response TokenResponse
type tokenResponse struct {
	// in:body
	content struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		Expiry       string `json:"expiry"`
	}
}
