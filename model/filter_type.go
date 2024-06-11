package model

// FilterType : FilterType model
type FilterType string

// List of possible FilterType values
const (
	FilterType_CROP                     FilterType = "CROP"
	FilterType_CONFORM                  FilterType = "CONFORM"
	FilterType_WATERMARK                FilterType = "WATERMARK"
	FilterType_ENHANCED_WATERMARK       FilterType = "ENHANCED_WATERMARK"
	FilterType_ROTATE                   FilterType = "ROTATE"
	FilterType_DEINTERLACE              FilterType = "DEINTERLACE"
	FilterType_ENHANCED_DEINTERLACE     FilterType = "ENHANCED_DEINTERLACE"
	FilterType_AUDIO_MIX                FilterType = "AUDIO_MIX"
	FilterType_DENOISE_HQDN3D           FilterType = "DENOISE_HQDN3D"
	FilterType_TEXT                     FilterType = "TEXT"
	FilterType_UNSHARP                  FilterType = "UNSHARP"
	FilterType_SCALE                    FilterType = "SCALE"
	FilterType_INTERLACE                FilterType = "INTERLACE"
	FilterType_AUDIO_VOLUME             FilterType = "AUDIO_VOLUME"
	FilterType_EBU_R128_SINGLE_PASS     FilterType = "EBU_R128_SINGLE_PASS"
	FilterType_AZURE_SPEECH_TO_CAPTIONS FilterType = "AZURE_SPEECH_TO_CAPTIONS"
)
