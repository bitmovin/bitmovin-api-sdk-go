package model

type ResponseEnvelope struct {
	// Unique correlation id
	RequestId string `json:"requestId,omitempty"`
	// Response status information
	Status ResponseStatus `json:"status,omitempty"`
	// Response information
	Data *ResultWrapper `json:"data,omitempty"`
	// Additional endpoint specific information
	More *map[string]interface{} `json:"more,omitempty"`
}

