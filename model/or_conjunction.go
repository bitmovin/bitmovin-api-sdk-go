package model

type OrConjunction struct {
	Type ConditionType `json:"type,omitempty"`
	// Array to perform the AND/OR evaluation on
	Conditions []AbstractCondition `json:"conditions,omitempty"`
}

