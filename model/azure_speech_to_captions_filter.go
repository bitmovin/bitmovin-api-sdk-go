package model

import (
	"encoding/json"
)

// AzureSpeechToCaptionsFilter model
type AzureSpeechToCaptionsFilter struct {
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
	CustomData                    *map[string]interface{}        `json:"customData,omitempty"`
	AzureSpeechToCaptionsSettings *AzureSpeechToCaptionsSettings `json:"azureSpeechToCaptionsSettings,omitempty"`
}

func (m AzureSpeechToCaptionsFilter) FilterType() FilterType {
	return FilterType_AZURE_SPEECH_TO_CAPTIONS
}
func (m AzureSpeechToCaptionsFilter) MarshalJSON() ([]byte, error) {
	type M AzureSpeechToCaptionsFilter
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AZURE_SPEECH_TO_CAPTIONS"

	return json.Marshal(x)
}
