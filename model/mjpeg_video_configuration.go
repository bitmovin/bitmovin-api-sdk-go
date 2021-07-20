package model

import (
	"encoding/json"
)

// MjpegVideoConfiguration model
type MjpegVideoConfiguration struct {
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
	// Width of the encoded video
	Width *int32 `json:"width,omitempty"`
	// Height of the encoded video
	Height *int32 `json:"height,omitempty"`
	// Target frame rate of the encoded video! (required)
	Rate *float64 `json:"rate,omitempty"`
	// The quality scale parameter (required)
	QScale      *int32      `json:"qScale,omitempty"`
	PixelFormat PixelFormat `json:"pixelFormat,omitempty"`
}

func (m MjpegVideoConfiguration) CodecConfigType() CodecConfigType {
	return CodecConfigType_MJPEG
}
func (m MjpegVideoConfiguration) MarshalJSON() ([]byte, error) {
	type M MjpegVideoConfiguration
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "MJPEG"

	return json.Marshal(x)
}
