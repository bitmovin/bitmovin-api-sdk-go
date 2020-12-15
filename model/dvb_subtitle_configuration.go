package model

import (
	"encoding/json"
)

// DvbSubtitleConfiguration model
type DvbSubtitleConfiguration struct {
	// Name of the resource. Can be freely chosen by the user. (required)
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
}

func (m DvbSubtitleConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_DVB_SUBTITLE
}
func (m DvbSubtitleConfiguration) MarshalJSON() ([]byte, error) {
	type M DvbSubtitleConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DVB_SUBTITLE"

	return json.Marshal(x)
}
