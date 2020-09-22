package model

import (
	"encoding/json"
)

// AdAnalyticsNotEqualFilter model
type AdAnalyticsNotEqualFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AdAnalyticsNotEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_NE
}
func (m AdAnalyticsNotEqualFilter) MarshalJSON() ([]byte, error) {
	type M AdAnalyticsNotEqualFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "NE"

	return json.Marshal(x)
}
