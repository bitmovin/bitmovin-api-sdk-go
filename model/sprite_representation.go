package model

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
