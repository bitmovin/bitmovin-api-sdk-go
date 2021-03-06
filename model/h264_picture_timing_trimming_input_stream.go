package model

import (
	"encoding/json"
)

// H264PictureTimingTrimmingInputStream model
type H264PictureTimingTrimmingInputStream struct {
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
	// The id of the ingest input stream that should be trimmed
	InputStreamId *string `json:"inputStreamId,omitempty"`
	// Defines the H264 SEI picture timing, as specified in ISO/IEC 14496-10:2008, of the frame from which the encoding should start. The frame indicated by this value will be included in the encoding
	StartPicTiming *string `json:"startPicTiming,omitempty"`
	// Defines the H264 SEI picture timing, as specified in ISO/IEC 14496-10:2008, of the frame at which the encoding should stop. The frame indicated by this value will be included in the encoding
	EndPicTiming *string `json:"endPicTiming,omitempty"`
}

func (m H264PictureTimingTrimmingInputStream) InputStreamType() InputStreamType {
	return InputStreamType_TRIMMING_H264_PICTURE_TIMING
}
func (m H264PictureTimingTrimmingInputStream) MarshalJSON() ([]byte, error) {
	type M H264PictureTimingTrimmingInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "TRIMMING_H264_PICTURE_TIMING"

	return json.Marshal(x)
}
