package model

// StreamsPublicSigningKey model
type StreamsPublicSigningKey struct {
	// The unique identifier of the key
	KeyId     *string   `json:"keyId,omitempty"`
	CreatedAt *DateTime `json:"createdAt,omitempty"`
}
