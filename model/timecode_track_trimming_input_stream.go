package model

import (
	"encoding/json"
)

// TimecodeTrackTrimmingInputStream model
type TimecodeTrackTrimmingInputStream struct {
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
	// The id of the ingest input stream that should be trimmed (required)
	InputStreamId *string `json:"inputStreamId,omitempty"`
	// Defines the timecode, in SMPTE-12M format, of the frame from which the encoding should start. The frame indicated by this value will be included in the encoding (required)
	StartTimeCode *string `json:"startTimeCode,omitempty"`
	// Defines the timecode, in SMPTE-12M format, of the frame at which the encoding should stop. The frame indicated by this value will be included in the encoding (required)
	EndTimeCode *string `json:"endTimeCode,omitempty"`
}

func (m TimecodeTrackTrimmingInputStream) InputStreamType() InputStreamType {
	return InputStreamType_TRIMMING_TIME_CODE_TRACK
}
func (m TimecodeTrackTrimmingInputStream) MarshalJSON() ([]byte, error) {
	type M TimecodeTrackTrimmingInputStream
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "TRIMMING_TIME_CODE_TRACK"

	return json.Marshal(x)
}
