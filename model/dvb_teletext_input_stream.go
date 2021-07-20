package model

import (
	"encoding/json"
)

// DvbTeletextInputStream model
type DvbTeletextInputStream struct {
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
	// Id of input
	InputId *string `json:"inputId,omitempty"`
	// Path to media file
	InputPath *string `json:"inputPath,omitempty"`
	// Specifies the algorithm how the stream in the input file will be selected
	SelectionMode StreamSelectionMode `json:"selectionMode,omitempty"`
	// Position of the stream, starting from 0.
	Position *int32 `json:"position,omitempty"`
	// Page number of the subtitles
	Page *int32 `json:"page,omitempty"`
}

func (m DvbTeletextInputStream) InputStreamType() InputStreamType {
	return InputStreamType_DVB_TELETEXT
}
func (m DvbTeletextInputStream) MarshalJSON() ([]byte, error) {
	type M DvbTeletextInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DVB_TELETEXT"

	return json.Marshal(x)
}
