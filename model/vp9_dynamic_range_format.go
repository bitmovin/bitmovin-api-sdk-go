package model

// Vp9DynamicRangeFormat : Configures what kind of dynamic range the output should conform to. Can be used to convert from SDR to HLG, from HLG to SDR.
type Vp9DynamicRangeFormat string

// List of possible Vp9DynamicRangeFormat values
const (
	Vp9DynamicRangeFormat_HLG Vp9DynamicRangeFormat = "HLG"
	Vp9DynamicRangeFormat_SDR Vp9DynamicRangeFormat = "SDR"
)
