package model

import (
	"bytes"
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
	// The CEA-608 caption channel to extract, as defined in ANSI/CTA-608-E. Only the primary channel of each field is selectable: CC1 on field 1 and CC3 on field 2. (required)
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

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
