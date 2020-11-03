package model

import (
    "encoding/json"
)

// AdAnalyticsNotContainsFilter model
type AdAnalyticsNotContainsFilter struct {
    Name *map[string]interface{} `json:"name,omitempty"`
    Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AdAnalyticsNotContainsFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_NOTCONTAINS
}
func (m AdAnalyticsNotContainsFilter) MarshalJSON() ([]byte, error) {
    type M AdAnalyticsNotContainsFilter
    x := struct {
        Operator string `json:"operator"`
        M
    }{M: M(m)}

    x.Operator = "NOTCONTAINS"

    return json.Marshal(x)
}


