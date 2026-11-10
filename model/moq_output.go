package model

import (
	"bytes"
	"encoding/json"
)

// Media over QUIC (MoQ) output. Publishes the live encoding to a MoQ relay over MoQ Transport (MoQT).  **BETA:** MoQ output is a beta feature. MoQ Transport is still an evolving IETF draft and relay implementations differ, so it might not work with some MoQ relays. It has been tested with `DRAFT_16` and the Cloudflare MoQ relay.
type MoqOutput struct {
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
	// URL of the MoQ relay the encoder publishes to. (required)
	RelayUrl *string `json:"relayUrl,omitempty"`
	// MoQT Track Namespace under which the tracks are published. A tuple of one to 32 fields, not a URI or path. (required)
	TrackNamespace []string `json:"trackNamespace,omitempty"`
	// MoQ Transport draft version to use for the session with the relay. (required)
	DraftVersion MoqDraftVersion `json:"draftVersion,omitempty"`
	// Token sent to the MoQ relay as the MoQT AUTHORIZATION TOKEN parameter. This is the only supported authentication mechanism; credentials in the relay URL are not used.
	AuthToken *string `json:"authToken,omitempty"`
	// Token type of the MoQT AUTHORIZATION TOKEN parameter (IANA MoQ Auth Token Type registry or private-use range). Only used together with authToken. If not set, the encoder applies its default token type.
	AuthTokenType *int32 `json:"authTokenType,omitempty"`
}

func (m MoqOutput) OutputType() OutputType {
	return OutputType_MOQ
}
func (m MoqOutput) MarshalJSON() ([]byte, error) {
	type M MoqOutput
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "MOQ"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
