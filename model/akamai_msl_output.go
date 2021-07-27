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
	Acl        []AclEntry              `json:"acl,omitempty"`
	// The Akamai stream ID (required)
	StreamId *int32 `json:"streamId,omitempty"`
	// The Akamai event name (required)
	EventName *string `json:"eventName,omitempty"`
	// The stream format. Only DASH is supported at the moment. (required)
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
