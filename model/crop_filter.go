package model

import (
	"encoding/json"
)

// CropFilter model
type CropFilter struct {
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
	Id *string `json:"id,omitempty"`
	// Amount of pixels which will be cropped of the input video from the left side.
	Left *int32 `json:"left,omitempty"`
	// Amount of pixels which will be cropped of the input video from the right side.
	Right *int32 `json:"right,omitempty"`
	// Amount of pixels which will be cropped of the input video from the top.
	Top *int32 `json:"top,omitempty"`
	// Amount of pixels which will be cropped of the input video from the bottom.
	Bottom *int32       `json:"bottom,omitempty"`
	Unit   PositionUnit `json:"unit,omitempty"`
}

func (m CropFilter) FilterType() FilterType {
	return FilterType_CROP
}
func (m CropFilter) MarshalJSON() ([]byte, error) {
	type M CropFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CROP"

	return json.Marshal(x)
}
