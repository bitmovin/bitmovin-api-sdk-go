package model

import (
	"bytes"
	"encoding/json"
)

// DashProgressiveWebmRepresentation model
type DashProgressiveWebmRepresentation struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// UUID of a muxing (required)
	MuxingId *string `json:"muxingId,omitempty"`
	// Used to signal a dependency with another representation. The representation may belong to a different adaptation set
	DependencyId *string `json:"dependencyId,omitempty"`
	// Path to the Progressive WebM file (required)
	FilePath *string `json:"filePath,omitempty"`
}

func (m DashProgressiveWebmRepresentation) DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator {
	return DashRepresentationTypeDiscriminator_PROGRESSIVE_WEBM
}
func (m DashProgressiveWebmRepresentation) MarshalJSON() ([]byte, error) {
	type M DashProgressiveWebmRepresentation
	x := struct {
		TypeDiscriminator string `json:"typeDiscriminator"`
		M
	}{M: M(m)}

	x.TypeDiscriminator = "PROGRESSIVE_WEBM"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
