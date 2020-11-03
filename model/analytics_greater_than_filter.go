package model

import (
    "encoding/json"
)

// AnalyticsGreaterThanFilter model
type AnalyticsGreaterThanFilter struct {
    Name *map[string]interface{} `json:"name,omitempty"`
    Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsGreaterThanFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_GT
}
func (m AnalyticsGreaterThanFilter) MarshalJSON() ([]byte, error) {
    type M AnalyticsGreaterThanFilter
    x := struct {
        Operator string `json:"operator"`
        M
    }{M: M(m)}

    x.Operator = "GT"

    return json.Marshal(x)
}


