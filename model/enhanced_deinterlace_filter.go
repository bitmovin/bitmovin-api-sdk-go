package model

import (
	"encoding/json"
)

// EnhancedDeinterlaceFilter model
type EnhancedDeinterlaceFilter struct {
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
	Id         *string                       `json:"id,omitempty"`
	Parity     EnhancedDeinterlaceParity     `json:"parity,omitempty"`
	Mode       EnhancedDeinterlaceMode       `json:"mode,omitempty"`
	AutoEnable EnhancedDeinterlaceAutoEnable `json:"autoEnable,omitempty"`
}

func (m EnhancedDeinterlaceFilter) FilterType() FilterType {
	return FilterType_ENHANCED_DEINTERLACE
}
func (m EnhancedDeinterlaceFilter) MarshalJSON() ([]byte, error) {
	type M EnhancedDeinterlaceFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "ENHANCED_DEINTERLACE"

	return json.Marshal(x)
}
