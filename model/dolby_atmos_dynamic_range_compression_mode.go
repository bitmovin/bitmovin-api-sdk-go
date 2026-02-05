package model

// DolbyAtmosDynamicRangeCompressionMode : Dynamic range compression processing mode
type DolbyAtmosDynamicRangeCompressionMode string

// List of possible DolbyAtmosDynamicRangeCompressionMode values
const (
	DolbyAtmosDynamicRangeCompressionMode_AUTO           DolbyAtmosDynamicRangeCompressionMode = "AUTO"
	DolbyAtmosDynamicRangeCompressionMode_NONE           DolbyAtmosDynamicRangeCompressionMode = "NONE"
	DolbyAtmosDynamicRangeCompressionMode_FILM_STANDARD  DolbyAtmosDynamicRangeCompressionMode = "FILM_STANDARD"
	DolbyAtmosDynamicRangeCompressionMode_FILM_LIGHT     DolbyAtmosDynamicRangeCompressionMode = "FILM_LIGHT"
	DolbyAtmosDynamicRangeCompressionMode_MUSIC_STANDARD DolbyAtmosDynamicRangeCompressionMode = "MUSIC_STANDARD"
	DolbyAtmosDynamicRangeCompressionMode_MUSIC_LIGHT    DolbyAtmosDynamicRangeCompressionMode = "MUSIC_LIGHT"
	DolbyAtmosDynamicRangeCompressionMode_SPEECH         DolbyAtmosDynamicRangeCompressionMode = "SPEECH"
)
