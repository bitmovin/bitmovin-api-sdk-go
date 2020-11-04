package model

import (
	"encoding/json"
)

// DeinterlaceFilter model
type DeinterlaceFilter struct {
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
	Id                 *string                       `json:"id,omitempty"`
	Parity             PictureFieldParity            `json:"parity,omitempty"`
	Mode               DeinterlaceMode               `json:"mode,omitempty"`
	FrameSelectionMode DeinterlaceFrameSelectionMode `json:"frameSelectionMode,omitempty"`
	AutoEnable         DeinterlaceAutoEnable         `json:"autoEnable,omitempty"`
}

func (m DeinterlaceFilter) FilterType() FilterType {
	return FilterType_DEINTERLACE
}
func (m DeinterlaceFilter) MarshalJSON() ([]byte, error) {
	type M DeinterlaceFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DEINTERLACE"

	return json.Marshal(x)
}
