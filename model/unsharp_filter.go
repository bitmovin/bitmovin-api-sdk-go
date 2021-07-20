package model

import (
	"encoding/json"
)

// UnsharpFilter model
type UnsharpFilter struct {
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
	// Must be an odd integer between 3 and 23
	LumaMatrixHorizontalSize *int32 `json:"lumaMatrixHorizontalSize,omitempty"`
	// Must be an odd integer between 3 and 23
	LumaMatrixVerticalSize *int32 `json:"lumaMatrixVerticalSize,omitempty"`
	// Negative value: blur, positive value: sharpen, floating point number, valid value range: -1.5 - 1.5
	LumaEffectStrength *float64 `json:"lumaEffectStrength,omitempty"`
	// Must be an odd integer between 3 and 23
	ChromaMatrixHorizontalSize *int32 `json:"chromaMatrixHorizontalSize,omitempty"`
	// Must be an odd integer between 3 and 23
	ChromaMatrixVerticalSize *int32 `json:"chromaMatrixVerticalSize,omitempty"`
	// Negative value: blur, positive value: sharpen, floating point number, valid value range: -1.5 - 1.5
	ChromaEffectStrength *float64 `json:"chromaEffectStrength,omitempty"`
}

func (m UnsharpFilter) FilterType() FilterType {
	return FilterType_UNSHARP
}
func (m UnsharpFilter) MarshalJSON() ([]byte, error) {
	type M UnsharpFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "UNSHARP"

	return json.Marshal(x)
}
