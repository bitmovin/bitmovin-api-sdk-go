package model

import (
	"encoding/json"
)

// SimpleEncodingLiveJobCdnOutput model
type SimpleEncodingLiveJobCdnOutput struct {
	// This sets the maximum output resolution that will be generated.
	MaxResolution SimpleEncodingLiveMaxResolution `json:"maxResolution,omitempty"`
}

func (m SimpleEncodingLiveJobCdnOutput) SimpleEncodingLiveJobOutputType() SimpleEncodingLiveJobOutputType {
	return SimpleEncodingLiveJobOutputType_CDN
}
func (m SimpleEncodingLiveJobCdnOutput) MarshalJSON() ([]byte, error) {
	type M SimpleEncodingLiveJobCdnOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CDN"

	return json.Marshal(x)
}
