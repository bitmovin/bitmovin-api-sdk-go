package model

import (
    "encoding/json"
)

// AnalyticsContainsFilter model
type AnalyticsContainsFilter struct {
    Name *map[string]interface{} `json:"name,omitempty"`
    Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsContainsFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_CONTAINS
}
func (m AnalyticsContainsFilter) MarshalJSON() ([]byte, error) {
    type M AnalyticsContainsFilter
    x := struct {
        Operator string `json:"operator"`
        M
    }{M: M(m)}

    x.Operator = "CONTAINS"

    return json.Marshal(x)
}


