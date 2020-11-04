package model

import (
	"encoding/json"
)

// AnalyticsEqualFilter model
type AnalyticsEqualFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_EQ
}
func (m AnalyticsEqualFilter) MarshalJSON() ([]byte, error) {
	type M AnalyticsEqualFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "EQ"

	return json.Marshal(x)
}
