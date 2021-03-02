package model

// InfrastructureSettings model
type InfrastructureSettings struct {
	// Id of a custom infrastructure, e.g., AWS Cloud Connect
	InfrastructureId *string     `json:"infrastructureId,omitempty"`
	CloudRegion      CloudRegion `json:"cloudRegion,omitempty"`
}
