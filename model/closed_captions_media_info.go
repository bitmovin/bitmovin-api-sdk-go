package model

type ClosedCaptionsMediaInfo struct {
	// The value is a quoted-string which specifies the group to which the Rendition belongs.
	GroupId string `json:"groupId,omitempty"`
	// Primary language in the rendition.
	Language string `json:"language,omitempty"`
	// Identifies a language that is associated with the Rendition.
	AssocLanguage string `json:"assocLanguage,omitempty"`
	// Human readable description of the rendition.
	Name string `json:"name,omitempty"`
	// If set to true, the client SHOULD play this Rendition of the content in the absence of information from the user.
	IsDefault *bool `json:"isDefault,omitempty"`
	// If set to true, the client MAY choose to play this Rendition in the absence of explicit user preference.
	Autoselect *bool `json:"autoselect,omitempty"`
	// Contains Uniform Type Identifiers
	Characteristics []string `json:"characteristics,omitempty"`
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Path to segments.
	SegmentPath string `json:"segmentPath,omitempty"`
	// Id of the encoding.
	EncodingId string `json:"encodingId,omitempty"`
	// Id of the stream.
	StreamId string `json:"streamId,omitempty"`
	// Id of the muxing.
	MuxingId string `json:"muxingId,omitempty"`
	// Id of the DRM.
	DrmId string `json:"drmId,omitempty"`
	// Number of the first segment. Default is 0.
	StartSegmentNumber *int64 `json:"startSegmentNumber,omitempty"`
	// Number of the last segment. Default is the last one that was encoded.
	EndSegmentNumber *int64 `json:"endSegmentNumber,omitempty"`
	// Specifies a Rendition within the segments in the Media Playlist. (See HLS spec 4.3.4.1. EXT-X-MEDIA INSTREAM-ID)
	InstreamId string `json:"instreamId,omitempty"`
	// A value of true indicates that the Rendition contains content which is considered essential to play.
	Forced *bool `json:"forced,omitempty"`
}

