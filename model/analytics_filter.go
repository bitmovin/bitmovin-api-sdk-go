package model

type AnalyticsFilter struct {
	Name string `json:"name,omitempty"`
	Operator AnalyticsOperator `json:"operator,omitempty"`
	// The value to compare to the property specified by the name (required)
	Value string `json:"value,omitempty"`
}

