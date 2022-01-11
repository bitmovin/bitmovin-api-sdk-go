package model

import (
	"encoding/json"
)

// DashImscRepresentation model
type DashImscRepresentation struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// URL of the referenced IMSC file (required)
	ImscUrl *string `json:"imscUrl,omitempty"`
}

func (m DashImscRepresentation) DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator {
	return DashRepresentationTypeDiscriminator_IMSC
}
func (m DashImscRepresentation) MarshalJSON() ([]byte, error) {
	type M DashImscRepresentation
	x := struct {
		TypeDiscriminator string `json:"typeDiscriminator"`
		M
	}{M: M(m)}

	x.TypeDiscriminator = "IMSC"

	return json.Marshal(x)
}
