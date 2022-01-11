package model

import (
	"encoding/json"
)

// DashVttRepresentation model
type DashVttRepresentation struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// URL of the referenced VTT file (required)
	VttUrl *string `json:"vttUrl,omitempty"`
}

func (m DashVttRepresentation) DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator {
	return DashRepresentationTypeDiscriminator_VTT
}
func (m DashVttRepresentation) MarshalJSON() ([]byte, error) {
	type M DashVttRepresentation
	x := struct {
		TypeDiscriminator string `json:"typeDiscriminator"`
		M
	}{M: M(m)}

	x.TypeDiscriminator = "VTT"

	return json.Marshal(x)
}
