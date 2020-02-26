package model

type DefaultManifestAttributeCondition struct {
	// The attribute that should be used for the evaluation: - audio.codec - audio.language - audio.bitrate - subtitle.format - subtitle.language - video.height - video.width - video.codec - video.bitrate - drm.type - muxing.type (required)
	Attribute string `json:"attribute,omitempty"`
	// The operator that should be used for the evaluation (required)
	Operator ConditionOperator `json:"operator,omitempty"`
	// The value that should be used for comparison. Valid values depend on the attribute: - audio.codec (Enum; e.g., AAC, MP3, OPUS) - audio.language (String) - audio.bitrate (Integer) - subtitle.format (Enum; e.g., WEBVTT) - subtitle.language (String) - video.height (Integer) - video.width (Integer) - video.codec (Enum; e.g., AV1, H264, VP9) - video.bitrate (Integer) - drm.type (Enum; NoDrm, Cenc, CencWidevine, CencPlayReady, CencMarlin, CencFairPlay, Aes128, ClearKey, PrimeTime, Widevine, PlayReady, Marlin, FairPlay) - muxing.type (Enum; e.g., FMP4, MP4) (required)
	Value string `json:"value,omitempty"`
}
func (o DefaultManifestAttributeCondition) ConditionType() ConditionType {
    return ConditionType_CONDITION
}

