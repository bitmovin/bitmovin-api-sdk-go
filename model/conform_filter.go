package model

import (
	"bytes"
	"encoding/json"
)

// ConformFilter model
type ConformFilter struct {
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
	// The FPS the input should be changed to.
	TargetFps *float64 `json:"targetFps,omitempty"`
}

func (m ConformFilter) FilterType() FilterType {
	return FilterType_CONFORM
}
func (m ConformFilter) MarshalJSON() ([]byte, error) {
	type M ConformFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CONFORM"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
