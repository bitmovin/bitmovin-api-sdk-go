package model

type ResponseError struct {
	// Unique correlation id
	RequestId string `json:"requestId,omitempty"`
	// Response status information
	Status ResponseStatus `json:"status,omitempty"`
	// Response information
	Data *ResponseErrorData `json:"data,omitempty"`
}

