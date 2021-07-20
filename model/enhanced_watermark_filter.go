package model

import (
	"encoding/json"
)

// EnhancedWatermarkFilter model
type EnhancedWatermarkFilter struct {
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
	// URL of the file to be used as watermark image. Supported image formats: PNG, JPEG, BMP, GIF (required)
	Image *string `json:"image,omitempty"`
	// Distance from the left edge of the input video to the left edge of the watermark image. May not be set if 'right' is set.
	Left *float64 `json:"left,omitempty"`
	// Distance from the right edge of the input video to the right edge of the watermark image . May not be set if 'left' is set.
	Right *float64 `json:"right,omitempty"`
	// Distance from the top edge of the input video to the top edge of the watermark image. May not be set if 'bottom' is set.
	Top *float64 `json:"top,omitempty"`
	// Distance from the bottom edge of the input video to the bottom edge of the watermark image. May not be set if 'top' is set.
	Bottom *float64     `json:"bottom,omitempty"`
	Unit   PositionUnit `json:"unit,omitempty"`
	// Opacity to apply on the watermark image. Valid values are from 0.0 (completely transparent) to 1.0 (not transparent at all)
	Opacity *float64 `json:"opacity,omitempty"`
	// Desired width of the watermark image, the unit of the parameter is specified separately by the parameter 'unit'. If both width and height are set the watermark size is fixed. If only one is set the aspect ratio of the image will be used to rescale it
	Width *float64 `json:"width,omitempty"`
	// Desired height of the watermark image, the unit of the parameter is specified separately by the parameter 'unit'. If both width and height are set the watermark size is fixed. If only one is set the aspect ratio of the image will be used to rescale it
	Height *float64 `json:"height,omitempty"`
}

func (m EnhancedWatermarkFilter) FilterType() FilterType {
	return FilterType_ENHANCED_WATERMARK
}
func (m EnhancedWatermarkFilter) MarshalJSON() ([]byte, error) {
	type M EnhancedWatermarkFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "ENHANCED_WATERMARK"

	return json.Marshal(x)
}
