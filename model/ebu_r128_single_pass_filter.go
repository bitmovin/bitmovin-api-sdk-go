package model

import (
	"encoding/json"
)

// EbuR128SinglePassFilter model
type EbuR128SinglePassFilter struct {
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
	// Set the targeted integrated loudness value. Range is from '-70.0' to '-5.0'. Default value is '-24.0'. Value is measured in LUFS (Loudness Units, referenced to Full Scale)
	IntegratedLoudness *float64 `json:"integratedLoudness,omitempty"`
	// Set the targeted loudness range target. Range is from '1.0' to '20.0'. Default value is '7.0'. Loudness range measures the variation of loudness on a macroscopic time-scale in units of LU (Loudness Units). For programmes shorter than 1 minute, the use of the measure Loudness Range is not recommended due to too few data points (Loudness Range is based on the Short-term-Loudness values (3-seconds-window)).
	LoudnessRange *float64 `json:"loudnessRange,omitempty"`
	// Set maximum true peak. Range is from '-9.0' to '0.0'. Default value is '-2.0'. Values are measured in dBTP (db True Peak)
	MaximumTruePeakLevel *float64 `json:"maximumTruePeakLevel,omitempty"`
}

func (m EbuR128SinglePassFilter) FilterType() FilterType {
	return FilterType_EBU_R128_SINGLE_PASS
}
func (m EbuR128SinglePassFilter) MarshalJSON() ([]byte, error) {
	type M EbuR128SinglePassFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "EBU_R128_SINGLE_PASS"

	return json.Marshal(x)
}
