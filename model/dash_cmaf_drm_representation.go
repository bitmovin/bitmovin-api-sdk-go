package model

type DashCmafDrmRepresentation struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// UUID of an encoding
	EncodingId string `json:"encodingId,omitempty"`
	// UUID of a muxing
	MuxingId string `json:"muxingId,omitempty"`
	Type DashMuxingType `json:"type,omitempty"`
	// Path to segments
	SegmentPath string `json:"segmentPath,omitempty"`
	// Number of the first segment
	StartSegmentNumber *int64 `json:"startSegmentNumber,omitempty"`
	// Number of the last segment. Default is the last one that was encoded
	EndSegmentNumber *int64 `json:"endSegmentNumber,omitempty"`
	// Id of the Keyframe to start with
	StartKeyframeId string `json:"startKeyframeId,omitempty"`
	// Id of the Keyframe to end with
	EndKeyframeId string `json:"endKeyframeId,omitempty"`
	// DRM Id
	DrmId string `json:"drmId,omitempty"`
}

