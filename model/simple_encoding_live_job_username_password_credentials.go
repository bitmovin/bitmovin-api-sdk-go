package model

// SimpleEncodingLiveJobUsernamePasswordCredentials model
type SimpleEncodingLiveJobUsernamePasswordCredentials struct {
	// The username to be used for authentication. (required)
	Username *string `json:"username,omitempty"`
	// The password to be used for authentication (required)
	Password *string `json:"password,omitempty"`
}
