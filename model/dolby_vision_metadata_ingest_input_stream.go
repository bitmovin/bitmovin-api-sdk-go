package model

import (
	"encoding/json"
)

// DolbyVisionMetadataIngestInputStream model
type DolbyVisionMetadataIngestInputStream struct {
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
	// Id of input (required)
	InputId *string `json:"inputId,omitempty"`
	// Path to Dolby Vision Metadata file (required)
	InputPath *string `json:"inputPath,omitempty"`
}

func (m DolbyVisionMetadataIngestInputStream) InputStreamType() InputStreamType {
	return InputStreamType_SIDECAR_DOLBY_VISION_METADATA
}
func (m DolbyVisionMetadataIngestInputStream) MarshalJSON() ([]byte, error) {
	type M DolbyVisionMetadataIngestInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "SIDECAR_DOLBY_VISION_METADATA"

	return json.Marshal(x)
}
