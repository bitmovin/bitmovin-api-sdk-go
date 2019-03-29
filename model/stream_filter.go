package model

type StreamFilter struct {
	// The id of the filter that should be used in the stream
	Id string `json:"id,omitempty"`
	// Defines the order in which filters are applied. Filters are applied in ascending order.
	Position *int32 `json:"position,omitempty"`
}

