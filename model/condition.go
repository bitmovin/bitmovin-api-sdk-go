package model

type Condition struct {
	Type ConditionType `json:"type,omitempty"`
	Attribute ConditionAttribute `json:"attribute,omitempty"`
	Operator ConditionOperator `json:"operator,omitempty"`
	// The value that should be used for comparison
	Value string `json:"value,omitempty"`
}

