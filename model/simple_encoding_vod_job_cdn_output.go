package model

import (
	"encoding/json"
)

// SimpleEncodingVodJobCdnOutput model
type SimpleEncodingVodJobCdnOutput struct {
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
