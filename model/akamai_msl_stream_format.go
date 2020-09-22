package model

// AkamaiMslStreamFormat : AkamaiMslStreamFormat model
type AkamaiMslStreamFormat string

// List of possible AkamaiMslStreamFormat values
const (
	AkamaiMslStreamFormat_DASH AkamaiMslStreamFormat = "DASH"
	AkamaiMslStreamFormat_HLS  AkamaiMslStreamFormat = "HLS"
	AkamaiMslStreamFormat_CMAF AkamaiMslStreamFormat = "CMAF"
)
