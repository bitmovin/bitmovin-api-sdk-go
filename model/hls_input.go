package model

import (
	"encoding/json"
)

// HlsInput model
type HlsInput struct {
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
	// URL of HLS input
	Url *string `json:"url,omitempty"`
	// Specifies the source for ad markers messages: - MANIFEST: Ad marker messages are read from tags in the HLS manifest - SEGMENTS: Ad marker messages are read from the content segments from the data stream
	AdMarkersSource AdMarkersSource `json:"adMarkersSource,omitempty"`
}

func (m HlsInput) InputType() InputType {
	return InputType_HLS
}
func (m HlsInput) MarshalJSON() ([]byte, error) {
	type M HlsInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "HLS"

	return json.Marshal(x)
}
