package model

type DefaultManifestOrCondition struct {
	// Array to perform the OR evaluation on. This conditions evaluates to true if at least one sub condition evaluates to true. 
	Conditions []DefaultManifestCondition `json:"conditions,omitempty"`
}
func (o DefaultManifestOrCondition) ConditionType() ConditionType {
    return ConditionType_OR
}

