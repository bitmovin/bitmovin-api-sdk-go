package model

import (
	"encoding/json"
)

// VideoAdaptationSet model
type VideoAdaptationSet struct {
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
}

func (m VideoAdaptationSet) AdaptationSetType() AdaptationSetType {
	return AdaptationSetType_VIDEO
}
func (m VideoAdaptationSet) MarshalJSON() ([]byte, error) {
	type M VideoAdaptationSet
	x := struct {
		Type string `json:"type"`
		M
	}{M: M(m)}

	x.Type = "VIDEO"

	return json.Marshal(x)
}
