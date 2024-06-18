package model

import (
	"encoding/json"
)

// DashFmp4DrmRepresentation model
type DashFmp4DrmRepresentation struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// UUID of a muxing (required)
	MuxingId *string `json:"muxingId,omitempty"`
	// Used to signal a dependency with another representation. The representation may belong to a different adaptation set
	DependencyId *string                    `json:"dependencyId,omitempty"`
	Type         DashRepresentationType     `json:"type,omitempty"`
	Mode         DashRepresentationTypeMode `json:"mode,omitempty"`
	// Path to segments. Will be used as the representation id if the type is set to TEMPLATE_ADAPTATION_SET (required)
	SegmentPath *string `json:"segmentPath,omitempty"`
	// Number of the first segment
	StartSegmentNumber *int64 `json:"startSegmentNumber,omitempty"`
	// Number of the last segment. Default is the last one that was encoded
	EndSegmentNumber *int64 `json:"endSegmentNumber,omitempty"`
	// Id of the keyframe to start with. It takes precedence over startSegmentNumber
	StartKeyframeId *string `json:"startKeyframeId,omitempty"`
	// Id of the keyframe to end with. It takes precedence over endSegmentNumber. The segment containing the end keyframe is not included in the representation.
	EndKeyframeId *string `json:"endKeyframeId,omitempty"`
	// DRM Id (required)
	DrmId *string `json:"drmId,omitempty"`
}

func (m DashFmp4DrmRepresentation) DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator {
	return DashRepresentationTypeDiscriminator_DRM_FMP4
}
func (m DashFmp4DrmRepresentation) MarshalJSON() ([]byte, error) {
	type M DashFmp4DrmRepresentation
	x := struct {
		TypeDiscriminator string `json:"typeDiscriminator"`
		M
	}{M: M(m)}

	x.TypeDiscriminator = "DRM_FMP4"

	return json.Marshal(x)
}
