package model

import (
	"encoding/json"
)

// TimeBasedTrimmingInputStream model
type TimeBasedTrimmingInputStream struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// The id of the ingest input stream that should be trimmed
	InputStreamId *string `json:"inputStreamId,omitempty"`
	// Defines the offset in seconds at which the encoding should start, beginning at 0. The frame indicated by this value will be included in the encoding
	Offset *float64 `json:"offset,omitempty"`
	// Defines how many seconds of the input will be encoded
	Duration *float64 `json:"duration,omitempty"`
}

func (m TimeBasedTrimmingInputStream) InputStreamType() InputStreamType {
	return InputStreamType_TRIMMING_TIME_BASED
}
func (m TimeBasedTrimmingInputStream) MarshalJSON() ([]byte, error) {
	type M TimeBasedTrimmingInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "TRIMMING_TIME_BASED"

	return json.Marshal(x)
}
