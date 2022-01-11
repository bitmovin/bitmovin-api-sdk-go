package model

import (
	"encoding/json"
)

// DashMp4Representation model
type DashMp4Representation struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// UUID of a muxing (required)
	MuxingId *string `json:"muxingId,omitempty"`
	// Used to signal a dependency with another representation. The representation may belong to a different adaptation set
	DependencyId *string `json:"dependencyId,omitempty"`
	// Path to the MP4 file (required)
	FilePath *string `json:"filePath,omitempty"`
	// The type of the dash representation
	Type DashOnDemandRepresentationType `json:"type,omitempty"`
}

func (m DashMp4Representation) DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator {
	return DashRepresentationTypeDiscriminator_MP4
}
func (m DashMp4Representation) MarshalJSON() ([]byte, error) {
	type M DashMp4Representation
	x := struct {
		TypeDiscriminator string `json:"typeDiscriminator"`
		M
	}{M: M(m)}

	x.TypeDiscriminator = "MP4"

	return json.Marshal(x)
}
