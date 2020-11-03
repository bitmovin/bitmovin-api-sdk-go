package model

import (
    "encoding/json"
)

// AnalyticsNotContainsFilter model
type AnalyticsNotContainsFilter struct {
    Name *map[string]interface{} `json:"name,omitempty"`
    Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsNotContainsFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_NOTCONTAINS
}
func (m AnalyticsNotContainsFilter) MarshalJSON() ([]byte, error) {
    type M AnalyticsNotContainsFilter
    x := struct {
        Operator string `json:"operator"`
        M
    }{M: M(m)}

    x.Operator = "NOTCONTAINS"

    return json.Marshal(x)
}


