package model

import (
	"bytes"
	"encoding/json"
)

// H265V2ConstantBitrateModeConfig model
type H265V2ConstantBitrateModeConfig struct {
	Type H265V2RateControlModeConfigType `json:"type,omitempty"`
	// Enable filler data for constant bitrate mode.
	Fillerdata *bool `json:"fillerdata,omitempty"`
}

func (m H265V2ConstantBitrateModeConfig) H265V2RateControlModeConfigType() H265V2RateControlModeConfigType {
	return H265V2RateControlModeConfigType_CONSTANT_BITRATE_MODE
}
func (m H265V2ConstantBitrateModeConfig) MarshalJSON() ([]byte, error) {
	type M H265V2ConstantBitrateModeConfig
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CONSTANT_BITRATE_MODE"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
