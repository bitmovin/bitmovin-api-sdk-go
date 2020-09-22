package model

import (
	"encoding/json"
)

// AnalyticsLessThanFilter model
type AnalyticsLessThanFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsLessThanFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_LT
}
func (m AnalyticsLessThanFilter) MarshalJSON() ([]byte, error) {
	type M AnalyticsLessThanFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "LT"

	return json.Marshal(x)
}
