package model

import (
	"encoding/json"
)

// AkamaiMslOutput model
type AkamaiMslOutput struct {
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
	// Deprecation notice: This property does not have any effect and will not be returned by GET endpoints
	Acl []AclEntry `json:"acl,omitempty"`
	// The Akamai stream ID (required)
	StreamId *int32 `json:"streamId,omitempty"`
	// The Akamai event name (required)
	EventName *string `json:"eventName,omitempty"`
	// - DASH: configure the Encoding with fMP4 or CMAF muxings and a DASH manifest. - HLS: configure the Encoding with TS muxings and an HLS manifest. - CMAF: configure the Encoding with fMP4 or CMAF muxings with both DASH and HLS manifests. (required)
	StreamFormat AkamaiMslStreamFormat `json:"streamFormat,omitempty"`
	// The Akamai MSL Version. Only MSL4 is supported at the moment. (required)
	MslVersion AkamaiMslVersion `json:"mslVersion,omitempty"`
}

func (m AkamaiMslOutput) OutputType() OutputType {
	return OutputType_AKAMAI_MSL
}
func (m AkamaiMslOutput) MarshalJSON() ([]byte, error) {
	type M AkamaiMslOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AKAMAI_MSL"

	return json.Marshal(x)
}
