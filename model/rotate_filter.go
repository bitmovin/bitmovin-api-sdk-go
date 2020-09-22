package model

import (
	"encoding/json"
)

// RotateFilter model
type RotateFilter struct {
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
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Rotation of the video in degrees. A positive value will rotate the video clockwise and a negative one counter clockwise. (required)
	Rotation *int32 `json:"rotation,omitempty"`
}

func (m RotateFilter) FilterType() FilterType {
	return FilterType_ROTATE
}
func (m RotateFilter) MarshalJSON() ([]byte, error) {
	type M RotateFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "ROTATE"

	return json.Marshal(x)
}
