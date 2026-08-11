package model

import (
	"bytes"
	"encoding/json"
)

// DolbyLoudnessFilter model
type DolbyLoudnessFilter struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Name of the resource. Can be freely chosen by the user.
	Name *string `json:"name,omitempty"`
	// Description of the resource. Can be freely chosen by the user.
	Description *string `json:"description,omitempty"`
	// Creation timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Modified timestamp, returned as UTC expressed in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ
	ModifiedAt *DateTime `json:"modifiedAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// The target integrated loudness the audio should be corrected to. Range is from '-31' to '-8'. Default value is '-24'. Value is measured in LKFS (Loudness, K-weighted, relative to Full Scale).
	TargetLoudness *int32 `json:"targetLoudness,omitempty"`
	// The maximum true-peak level the corrected audio may reach. Range is from '-8.0' to '-0.1'. Default value is '-2.0'. Values are measured in dBTP (dB True Peak). Note that the maximum true peak level must be set at least 6 dB above the target loudness.
	MaximumTruePeakLevel *float64 `json:"maximumTruePeakLevel,omitempty"`
	// Whether to use the Dolby Dialogue Intelligence feature, which identifies and analyzes dialogue segments within the audio as a basis for speech gating. Default value is 'ENABLED'.
	DialogueIntelligence DolbyLoudnessDialogueIntelligence `json:"dialogueIntelligence,omitempty"`
	// The percentage of speech that must be detected within the audio before the dialogue loudness is used as the basis for loudness correction. Range is from '0' to '100'. Default value is '20'. This is only applied when dialogueIntelligence is 'ENABLED', as it selects between speech-gated and un-gated loudness measurement.
	SpeechDetectionThreshold *int32 `json:"speechDetectionThreshold,omitempty"`
	// The form of the content, used to optimize the loudness measurement gating. Content longer than 3 minutes (180 seconds) is considered long-form, shorter content is considered short-form. Default value is 'AUTO_DETECT'.
	ContentForm DolbyLoudnessContentForm `json:"contentForm,omitempty"`
}

func (m DolbyLoudnessFilter) FilterType() FilterType {
	return FilterType_DOLBY_LOUDNESS
}
func (m DolbyLoudnessFilter) MarshalJSON() ([]byte, error) {
	type M DolbyLoudnessFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "DOLBY_LOUDNESS"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
