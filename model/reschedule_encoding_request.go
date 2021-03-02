package model

// RescheduleEncodingRequest model
type RescheduleEncodingRequest struct {
	// Id of a custom infrastructure, e.g., AWS Cloud Connect
	InfrastructureId *string `json:"infrastructureId,omitempty"`
}
