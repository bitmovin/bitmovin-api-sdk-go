package model

import (
	"bytes"
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

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
