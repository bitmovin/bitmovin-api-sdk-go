package model

// HeAacV2Signaling : HeAacV2Signaling model
type HeAacV2Signaling string

// List of possible HeAacV2Signaling values
const (
	HeAacV2Signaling_IMPLICIT              HeAacV2Signaling = "IMPLICIT"
	HeAacV2Signaling_EXPLICIT_SBR          HeAacV2Signaling = "EXPLICIT_SBR"
	HeAacV2Signaling_EXPLICIT_PS           HeAacV2Signaling = "EXPLICIT_PS"
	HeAacV2Signaling_EXPLICIT_HIERARCHICAL HeAacV2Signaling = "EXPLICIT_HIERARCHICAL"
)
