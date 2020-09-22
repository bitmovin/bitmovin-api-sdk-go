package model

import (
	"encoding/json"
)

// AdAnalyticsLessThanFilter model
type AdAnalyticsLessThanFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AdAnalyticsLessThanFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_LT
}
func (m AdAnalyticsLessThanFilter) MarshalJSON() ([]byte, error) {
	type M AdAnalyticsLessThanFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "LT"

	return json.Marshal(x)
}
