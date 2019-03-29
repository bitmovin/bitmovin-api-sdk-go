package model

type StreamCondition struct {
	Type ConditionType `json:"type,omitempty"`
	Attribute StreamConditionAttribute `json:"attribute,omitempty"`
	Operator ConditionOperator `json:"operator,omitempty"`
	// The value that should be used for comparison
	Value string `json:"value,omitempty"`
}

