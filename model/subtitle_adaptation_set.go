package model

import (
	"encoding/json"
)

// SubtitleAdaptationSet model
type SubtitleAdaptationSet struct {
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
	// ISO 639-1 (Alpha-2) code identifying the language of the subtitle adaptation set (required)
	Lang *string `json:"lang,omitempty"`
}

func (m SubtitleAdaptationSet) AdaptationSetType() AdaptationSetType {
	return AdaptationSetType_SUBTITLE
}
func (m SubtitleAdaptationSet) MarshalJSON() ([]byte, error) {
	type M SubtitleAdaptationSet
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "SUBTITLE"

	return json.Marshal(x)
}
