package model

// ESAM condition following the SCTE-250 standard with timing offset and directional marker (OUT/IN) for signaling content boundaries
type EsamCondition struct {
	// The offset from the matched signal when this condition applies in ISO 8601 duration format, accurate to milliseconds (required)
	Offset *string `json:"offset,omitempty"`
	// Direction marker indicating the boundary type (OUT for start, IN for end) (required)
	Direction EsamDirection `json:"direction,omitempty"`
}
