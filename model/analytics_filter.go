package model

type AnalyticsFilter struct {
	Name string `json:"name,omitempty"`
	Operator AnalyticsOperator `json:"operator,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

