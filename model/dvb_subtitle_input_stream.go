package model

import (
	"encoding/json"
)

// DvbSubtitleInputStream model
type DvbSubtitleInputStream struct {
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
	// ID of an Input resource defining the input storage
	InputId *string `json:"inputId,omitempty"`
	// Path to an input media file
	InputPath *string `json:"inputPath,omitempty"`
	// Specifies the strategy for selecting a stream from the input file. Supported values for VOD encodings: AUTO, POSITION_ABSOLUTE, SUBTITLE_RELATIVE. Supported values for LIVE encodings: POSITION_ABSOLUTE
	SelectionMode StreamSelectionMode `json:"selectionMode,omitempty"`
	// Position of the stream to be selected from the input file (zero-based). Must not be set in combination with selectionMode 'AUTO', defaults to 0 for any other selectionMode.
	Position *int32 `json:"position,omitempty"`
}

func (m DvbSubtitleInputStream) InputStreamType() InputStreamType {
	return InputStreamType_DVB_SUBTITLE
}
func (m DvbSubtitleInputStream) MarshalJSON() ([]byte, error) {
	type M DvbSubtitleInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DVB_SUBTITLE"

	return json.Marshal(x)
}
