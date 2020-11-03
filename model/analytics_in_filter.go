package model

import (
    "encoding/json"
)

// AnalyticsInFilter model
type AnalyticsInFilter struct {
    Name *map[string]interface{} `json:"name,omitempty"`
    Value []map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsInFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_IN
}
func (m AnalyticsInFilter) MarshalJSON() ([]byte, error) {
    type M AnalyticsInFilter
    x := struct {
        Operator string `json:"operator"`
        M
    }{M: M(m)}

    x.Operator = "IN"

    return json.Marshal(x)
}


