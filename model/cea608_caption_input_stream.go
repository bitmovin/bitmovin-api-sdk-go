package model

import (
	"encoding/json"
)

// Cea608CaptionInputStream model
type Cea608CaptionInputStream struct {
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
	// Id of the Input (required)
	InputId *string `json:"inputId,omitempty"`
	// Path to media file (required)
	InputPath *string `json:"inputPath,omitempty"`
	// The channel number of the subtitle on the respective stream position (required)
	Channel Cea608ChannelType `json:"channel,omitempty"`
}

func (m Cea608CaptionInputStream) InputStreamType() InputStreamType {
	return InputStreamType_CAPTION_CEA608
}
func (m Cea608CaptionInputStream) MarshalJSON() ([]byte, error) {
	type M Cea608CaptionInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CAPTION_CEA608"

	return json.Marshal(x)
}
