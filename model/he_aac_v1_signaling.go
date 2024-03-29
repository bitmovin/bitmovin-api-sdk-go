package model

// HeAacV1Signaling : HeAacV1Signaling model
type HeAacV1Signaling string

// List of possible HeAacV1Signaling values
const (
	HeAacV1Signaling_IMPLICIT              HeAacV1Signaling = "IMPLICIT"
	HeAacV1Signaling_EXPLICIT_SBR          HeAacV1Signaling = "EXPLICIT_SBR"
	HeAacV1Signaling_EXPLICIT_HIERARCHICAL HeAacV1Signaling = "EXPLICIT_HIERARCHICAL"
)
