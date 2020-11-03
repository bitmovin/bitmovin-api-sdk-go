package model

import (
    "encoding/json"
)

// AnalyticsLessThanOrEqualFilter model
type AnalyticsLessThanOrEqualFilter struct {
    Name *map[string]interface{} `json:"name,omitempty"`
    Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsLessThanOrEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_LTE
}
func (m AnalyticsLessThanOrEqualFilter) MarshalJSON() ([]byte, error) {
    type M AnalyticsLessThanOrEqualFilter
    x := struct {
        Operator string `json:"operator"`
        M
    }{M: M(m)}

    x.Operator = "LTE"

    return json.Marshal(x)
}


