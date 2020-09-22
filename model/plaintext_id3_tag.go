package model

import (
	"encoding/json"
)

// PlaintextId3Tag model
type PlaintextId3Tag struct {
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
	Id           *string            `json:"id,omitempty"`
	PositionMode Id3TagPositionMode `json:"positionMode,omitempty"`
	// Frame number at which the Tag should be inserted
	Frame *int64 `json:"frame,omitempty"`
	// Time in seconds where the Tag should be inserted
	Time *float64 `json:"time,omitempty"`
	// Plain Text Data (required)
	Text *string `json:"text,omitempty"`
	// 4 character long Frame ID (required)
	FrameId *string `json:"frameId,omitempty"`
}

func (m PlaintextId3Tag) Id3TagType() Id3TagType {
	return Id3TagType_PLAIN_TEXT
}
func (m PlaintextId3Tag) MarshalJSON() ([]byte, error) {
	type M PlaintextId3Tag
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "PLAIN_TEXT"

	return json.Marshal(x)
}
