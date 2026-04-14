package model

import (
	"bytes"
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

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
