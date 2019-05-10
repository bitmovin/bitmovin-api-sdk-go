package model

type Condition struct {
	// The attribute that should be used for the evaluation
	Attribute string `json:"attribute,omitempty"`
	// The operator that should be used for the evaluation
	Operator ConditionOperator `json:"operator,omitempty"`
	// The value that should be used for comparison
	Value string `json:"value,omitempty"`
}
func (o Condition) ConditionType() ConditionType {
    return ConditionType_CONDITION
}

