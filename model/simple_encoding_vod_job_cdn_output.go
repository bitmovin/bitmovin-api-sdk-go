package model

import (
	"encoding/json"
)

// SimpleEncodingVodJobCdnOutput model
type SimpleEncodingVodJobCdnOutput struct {
	// List of artifacts created by the encoding job. Artifacts are files essential for playback of the generated content, e.g. manifests.
	Artifacts []SimpleEncodingVodJobOutputArtifact `json:"artifacts,omitempty"`
	// This sets the maximum output resolution that will be generated.
	MaxResolution SimpleEncodingLiveMaxResolution `json:"maxResolution,omitempty"`
}

func (m SimpleEncodingVodJobCdnOutput) SimpleEncodingVodJobOutputType() SimpleEncodingVodJobOutputType {
	return SimpleEncodingVodJobOutputType_CDN
}
func (m SimpleEncodingVodJobCdnOutput) MarshalJSON() ([]byte, error) {
	type M SimpleEncodingVodJobCdnOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "CDN"

	return json.Marshal(x)
}
