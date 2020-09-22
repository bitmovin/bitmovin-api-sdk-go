package model

// AudioVolumeFormat : Audio volume format
type AudioVolumeFormat string

// List of possible AudioVolumeFormat values
const (
	AudioVolumeFormat_U8   AudioVolumeFormat = "U8"
	AudioVolumeFormat_S16  AudioVolumeFormat = "S16"
	AudioVolumeFormat_S32  AudioVolumeFormat = "S32"
	AudioVolumeFormat_U8P  AudioVolumeFormat = "U8P"
	AudioVolumeFormat_S16P AudioVolumeFormat = "S16P"
	AudioVolumeFormat_S32P AudioVolumeFormat = "S32P"
	AudioVolumeFormat_S64  AudioVolumeFormat = "S64"
	AudioVolumeFormat_S64P AudioVolumeFormat = "S64P"
	AudioVolumeFormat_FLT  AudioVolumeFormat = "FLT"
	AudioVolumeFormat_FLTP AudioVolumeFormat = "FLTP"
	AudioVolumeFormat_NONE AudioVolumeFormat = "NONE"
	AudioVolumeFormat_DBL  AudioVolumeFormat = "DBL"
	AudioVolumeFormat_DBLP AudioVolumeFormat = "DBLP"
)
