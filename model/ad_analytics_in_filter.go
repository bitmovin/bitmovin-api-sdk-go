package model

import (
	"encoding/json"
)

// AdAnalyticsInFilter model
type AdAnalyticsInFilter struct {
	Name  *map[string]interface{}  `json:"name,omitempty"`
	Value []map[string]interface{} `json:"value,omitempty"`
}

func (m AdAnalyticsInFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_IN
}
func (m AdAnalyticsInFilter) MarshalJSON() ([]byte, error) {
	type M AdAnalyticsInFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "IN"

	return json.Marshal(x)
}
