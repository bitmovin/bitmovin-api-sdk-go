package model

// DashRepresentationTypeResponse model
type DashRepresentationTypeResponse struct {
	// The type of the DASH representation
	Type DashRepresentationTypeDiscriminator `json:"type,omitempty"`
}
