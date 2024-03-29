package model

import (
	"encoding/json"
)

// AudioAdaptationSet model
type AudioAdaptationSet struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Custom adaptation set attributes
	CustomAttributes []CustomAttribute `json:"customAttributes,omitempty"`
	// Roles of the adaptation set
	Roles []AdaptationSetRole `json:"roles,omitempty"`
	// Provide signaling of CEA 607 and CEA 708
	Accessibilities []Accessibility `json:"accessibilities,omitempty"`
	// List of labels
	Labels []Label `json:"labels,omitempty"`
	// ISO 639-1 (Alpha-2) code identifying the language of the audio adaptation set (required)
	Lang *string `json:"lang,omitempty"`
}

func (m AudioAdaptationSet) AdaptationSetType() AdaptationSetType {
	return AdaptationSetType_AUDIO
}
func (m AudioAdaptationSet) MarshalJSON() ([]byte, error) {
	type M AudioAdaptationSet
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "AUDIO"

	return json.Marshal(x)
}
