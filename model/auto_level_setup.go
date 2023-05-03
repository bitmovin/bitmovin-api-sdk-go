package model

// AutoLevelSetup : Enable/disable automatic calculation of level, maxBitrate, and bufsize based on the least level that satisfies maximum property values for picture resolution, frame rate, and bit rate.
type AutoLevelSetup string

// List of possible AutoLevelSetup values
const (
	AutoLevelSetup_ENABLED  AutoLevelSetup = "ENABLED"
	AutoLevelSetup_DISABLED AutoLevelSetup = "DISABLED"
)
