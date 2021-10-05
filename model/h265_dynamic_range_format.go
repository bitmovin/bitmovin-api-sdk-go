package model

// H265DynamicRangeFormat : Configures what kind of dynamic range the output should conform to. Can be used to convert from SDR to HDR, from HDR to SDR or between different HDR formats
type H265DynamicRangeFormat string

// List of possible H265DynamicRangeFormat values
const (
	H265DynamicRangeFormat_DOLBY_VISION H265DynamicRangeFormat = "DOLBY_VISION"
	H265DynamicRangeFormat_HDR10        H265DynamicRangeFormat = "HDR10"
	H265DynamicRangeFormat_HLG          H265DynamicRangeFormat = "HLG"
	H265DynamicRangeFormat_SDR          H265DynamicRangeFormat = "SDR"
)
