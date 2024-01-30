package model

import (
	"encoding/json"
)

// RedundantRtmpInput model
type RedundantRtmpInput struct {
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
	// When there is no input signal present and this threshold in seconds is reached it will switch to another ingest point
	DelayThreshold *int32 `json:"delayThreshold,omitempty"`
	// Configuration for ingest points that use a dynamic IP based endpoint to stream to e.g.: rtmp://41.167.11.21/live Either ingestPoints **or** staticIngestPoints can be set
	IngestPoints []RtmpIngestPoint `json:"ingestPoints,omitempty"`
	// Configuration for static ingest points. These ingest points use a consistent endpoint to stream to e.g.: rtmps://live-ingest.bitmovin.com/live Either ingestPoints **or** staticIngestPoints can be set
	StaticIngestPoints []StaticRtmpIngestPoint `json:"staticIngestPoints,omitempty"`
}

func (m RedundantRtmpInput) InputType() InputType {
	return InputType_REDUNDANT_RTMP
}
func (m RedundantRtmpInput) MarshalJSON() ([]byte, error) {
	type M RedundantRtmpInput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "REDUNDANT_RTMP"

	return json.Marshal(x)
}
