package model

// StreamsSigningKeyResponse model
type StreamsSigningKeyResponse struct {
	// base64-encoded PEM file of the private key
	PrivateKey *string `json:"privateKey,omitempty"`
	// The unique identifier of the key
	KeyId   *string `json:"keyId,omitempty"`
	Message *string `json:"message,omitempty"`
}
