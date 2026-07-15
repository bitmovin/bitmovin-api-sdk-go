package model

import (
	"bytes"
	"encoding/json"
)

// H265V2PerceptualQualityModeConfig model
type H265V2PerceptualQualityModeConfig struct {
	Type H265V2RateControlModeConfigType `json:"type,omitempty"`
	// HVS-based perceptual encoding mode.
	PerceptualEncodingMode H265V2PerceptualEncodingMode `json:"perceptualEncodingMode,omitempty"`
	// Overall strength of perceptual model (0.0-1.0).
	PerceptualStrength *float64 `json:"perceptualStrength,omitempty"`
	// Strength of penalties from perceptual model (0.0-1.0).
	PerceptualPenaltyStrength *float64 `json:"perceptualPenaltyStrength,omitempty"`
	// Knee point of penalty strength modulation (0.0-1.0).
	PerceptualPenaltyKnee *float64 `json:"perceptualPenaltyKnee,omitempty"`
	// Strength of temporal component of perceptual model (0.0-1.0).
	PerceptualTemporalStrength *float64 `json:"perceptualTemporalStrength,omitempty"`
	// Pixels per degree (horizontal), i.e. width / FOV.
	PixelPerDegree *float64 `json:"pixelPerDegree,omitempty"`
}

func (m H265V2PerceptualQualityModeConfig) H265V2RateControlModeConfigType() H265V2RateControlModeConfigType {
	return H265V2RateControlModeConfigType_PERCEPTUAL_QUALITY_MODE
}
func (m H265V2PerceptualQualityModeConfig) MarshalJSON() ([]byte, error) {
	type M H265V2PerceptualQualityModeConfig
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "PERCEPTUAL_QUALITY_MODE"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
