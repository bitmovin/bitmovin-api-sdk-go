package model

import (
    "encoding/json"
)

// WatermarkFilter model
type WatermarkFilter struct {
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
    // URL of the file to be used as watermark image. Supported image formats: PNG, JPEG, BMP, GIF (required)
    Image *string `json:"image,omitempty"`
    // Distance from the left edge of the input video to the left edge of the watermark image. May not be set if 'right' is set.
    Left *int32 `json:"left,omitempty"`
    // Distance from the right edge of the input video to the right edge of the watermark image . May not be set if 'left' is set.
    Right *int32 `json:"right,omitempty"`
    // Distance from the top edge of the input video to the top edge of the watermark image. May not be set if 'bottom' is set.
    Top *int32 `json:"top,omitempty"`
    // Distance from the bottom edge of the input video to the bottom edge of the watermark image. May not be set if 'top' is set.
    Bottom *int32 `json:"bottom,omitempty"`
    Unit PositionUnit `json:"unit,omitempty"`
}

func (m WatermarkFilter) FilterType() FilterType {
    return FilterType_WATERMARK
}
func (m WatermarkFilter) MarshalJSON() ([]byte, error) {
    type M WatermarkFilter
    x := struct {
        Type string `json:"type"`
        M
    }{M: M(m)}

    x.Type = "WATERMARK"

    return json.Marshal(x)
}


