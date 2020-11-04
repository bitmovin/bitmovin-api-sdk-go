package model

// H264Trellis : Enables or disables Trellis quantization. NOTE: This requires cabac
type H264Trellis string

// List of possible H264Trellis values
const (
	H264Trellis_DISABLED         H264Trellis = "DISABLED"
	H264Trellis_ENABLED_FINAL_MB H264Trellis = "ENABLED_FINAL_MB"
	H264Trellis_ENABLED_ALL      H264Trellis = "ENABLED_ALL"
)
