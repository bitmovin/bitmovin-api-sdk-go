package model

// Av1DynamicRangeFormat : Configures what kind of dynamic range the output should conform to. Can be used to convert between different HDR formats.
type Av1DynamicRangeFormat string

// List of possible Av1DynamicRangeFormat values
const (
	Av1DynamicRangeFormat_DOLBY_VISION_PROFILE_10_0 Av1DynamicRangeFormat = "DOLBY_VISION_PROFILE_10_0"
	Av1DynamicRangeFormat_DOLBY_VISION_PROFILE_10_1 Av1DynamicRangeFormat = "DOLBY_VISION_PROFILE_10_1"
	Av1DynamicRangeFormat_HDR10                     Av1DynamicRangeFormat = "HDR10"
	Av1DynamicRangeFormat_SDR                       Av1DynamicRangeFormat = "SDR"
)
