package model

import (
	"encoding/json"
)

// IngestInputStream model
type IngestInputStream struct {
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
	// Id of input
	InputId *string `json:"inputId,omitempty"`
	// Path to media file
	InputPath *string `json:"inputPath,omitempty"`
	// Specifies the algorithm how the stream in the input file will be selected
	SelectionMode StreamSelectionMode `json:"selectionMode,omitempty"`
	// Position of the stream
	Position *int32 `json:"position,omitempty"`
}

func (m IngestInputStream) InputStreamType() InputStreamType {
	return InputStreamType_INGEST
}
func (m IngestInputStream) MarshalJSON() ([]byte, error) {
	type M IngestInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "INGEST"

	return json.Marshal(x)
}
