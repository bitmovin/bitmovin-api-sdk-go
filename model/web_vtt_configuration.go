package model

import (
	"encoding/json"
)

// WebVttConfiguration model
type WebVttConfiguration struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
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
	// If set to true, the hours section on webvtt timestamp values will explicitely have zeroes instead of being omitted for values where hours = 0.
	AppendOptionalZeroHour *bool `json:"appendOptionalZeroHour,omitempty"`
	// If set to true, the region information of the resulting webvtt file will be omitted. Defaults to false.
	IgnoreRegion        *bool                     `json:"ignoreRegion,omitempty"`
	CueIdentifierPolicy WebVttCueIdentifierPolicy `json:"cueIdentifierPolicy,omitempty"`
	Styling             *WebVttStyling            `json:"styling,omitempty"`
}

func (m WebVttConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_WEBVTT
}
func (m WebVttConfiguration) MarshalJSON() ([]byte, error) {
	type M WebVttConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "WEBVTT"

	return json.Marshal(x)
}
