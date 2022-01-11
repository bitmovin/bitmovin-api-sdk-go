package model

import (
	"encoding/json"
)

// ContentProtection model
type ContentProtection struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// UUID of a muxing (required)
	MuxingId *string `json:"muxingId,omitempty"`
	// Used to signal a dependency with another representation. The representation may belong to a different adaptation set
	DependencyId *string `json:"dependencyId,omitempty"`
	// DRM Id (required)
	DrmId *string `json:"drmId,omitempty"`
}

func (m ContentProtection) DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator {
	return DashRepresentationTypeDiscriminator_CONTENT_PROTECTION
}
func (m ContentProtection) MarshalJSON() ([]byte, error) {
	type M ContentProtection
	x := struct {
		TypeDiscriminator string `json:"typeDiscriminator"`
		M
	}{M: M(m)}

	x.TypeDiscriminator = "CONTENT_PROTECTION"

	return json.Marshal(x)
}
