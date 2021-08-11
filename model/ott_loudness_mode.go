package model

// OttLoudnessMode : Specifies the OTT loudness mode. If the mode is `DTSX_OTT_LOUDNESS_DETECT`, the `lkfs` value will be calculated internally and does not need to be provided. A provided `lkfs` value will be ignored. For the modes `DTSX_OTT_LOUDNESS_INPUT` and `DTSX_OTT_LOUDNESS_TARGET` an `lkfs` value must be provided by the user.
type OttLoudnessMode string

// List of possible OttLoudnessMode values
const (
	OttLoudnessMode_DTSX_OTT_LOUDNESS_DETECT OttLoudnessMode = "DTSX_OTT_LOUDNESS_DETECT"
	OttLoudnessMode_DTSX_OTT_LOUDNESS_INPUT  OttLoudnessMode = "DTSX_OTT_LOUDNESS_INPUT"
	OttLoudnessMode_DTSX_OTT_LOUDNESS_TARGET OttLoudnessMode = "DTSX_OTT_LOUDNESS_TARGET"
)
