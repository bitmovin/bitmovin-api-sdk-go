package model

type DefaultManifestAndCondition struct {
	// Array to perform the AND evaluation on. This conditions evaluates to true if all sub conditions evaluate to true. 
	Conditions []DefaultManifestCondition `json:"conditions,omitempty"`
}
func (o DefaultManifestAndCondition) ConditionType() ConditionType {
    return ConditionType_AND
}

