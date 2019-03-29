package model
type PresetConfiguration string

// List of PresetConfiguration
const (
	PresetConfiguration_LIVE_HIGH_QUALITY PresetConfiguration = "LIVE_HIGH_QUALITY"
	PresetConfiguration_LIVE_LOW_LATENCY PresetConfiguration = "LIVE_LOW_LATENCY"
	PresetConfiguration_VOD_HIGH_QUALITY PresetConfiguration = "VOD_HIGH_QUALITY"
	PresetConfiguration_VOD_HIGH_SPEED PresetConfiguration = "VOD_HIGH_SPEED"
)