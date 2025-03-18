package model

// PoisEndpointCredentials model
type PoisEndpointCredentials struct {
	// The username required to authenticate with the POIS server.
	Username *string `json:"username,omitempty"`
	// The password required for authentication with the POIS server.
	Password *string `json:"password,omitempty"`
}
