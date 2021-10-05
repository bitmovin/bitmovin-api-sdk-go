package model

import (
	"encoding/json"
)

// DolbyVisionInputStream model
type DolbyVisionInputStream struct {
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
	// Id of input (required)
	InputId *string `json:"inputId,omitempty"`
	// Path to Dolby Vision input video file. (required)
	VideoInputPath *string `json:"videoInputPath,omitempty"`
	// Path to Dolby Vision Metadata file. This field is required when the metadata is not embedded in the video input file.
	MetadataInputPath *string `json:"metadataInputPath,omitempty"`
}

func (m DolbyVisionInputStream) InputStreamType() InputStreamType {
	return InputStreamType_DOLBY_VISION
}
func (m DolbyVisionInputStream) MarshalJSON() ([]byte, error) {
	type M DolbyVisionInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DOLBY_VISION"

	return json.Marshal(x)
}
