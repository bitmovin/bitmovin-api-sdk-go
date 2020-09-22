package model

import (
	"encoding/json"
)

// FrameIdId3Tag model
type FrameIdId3Tag struct {
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
	// Base64 Encoded Data (required)
	Bytes *string `json:"bytes,omitempty"`
	// 4 character long Frame ID (required)
	FrameId *string `json:"frameId,omitempty"`
}

func (m FrameIdId3Tag) Id3TagType() Id3TagType {
	return Id3TagType_FRAME_ID
}
func (m FrameIdId3Tag) MarshalJSON() ([]byte, error) {
	type M FrameIdId3Tag
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "FRAME_ID"

	return json.Marshal(x)
}
