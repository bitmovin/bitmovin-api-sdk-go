package model

import (
	"encoding/json"
)

// SpriteRepresentation model
type SpriteRepresentation struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// UUID of an encoding (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// UUID of a stream (required)
	StreamId *string `json:"streamId,omitempty"`
	// UUID of a Sprite (required)
	SpriteId *string `json:"spriteId,omitempty"`
	// Path to sprite segments. Will be used as the representation id in the manifest. (required)
	SegmentPath *string `json:"segmentPath,omitempty"`
}

func (m SpriteRepresentation) DashRepresentationTypeDiscriminator() DashRepresentationTypeDiscriminator {
	return DashRepresentationTypeDiscriminator_SPRITE
}
func (m SpriteRepresentation) MarshalJSON() ([]byte, error) {
	type M SpriteRepresentation
	x := struct {
		TypeDiscriminator string `json:"typeDiscriminator"`
		M
	}{M: M(m)}

	x.TypeDiscriminator = "SPRITE"

	return json.Marshal(x)
}
