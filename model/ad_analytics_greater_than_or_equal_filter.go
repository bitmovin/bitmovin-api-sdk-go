package model

import (
    "encoding/json"
)

// AdAnalyticsGreaterThanOrEqualFilter model
type AdAnalyticsGreaterThanOrEqualFilter struct {
    Name *map[string]interface{} `json:"name,omitempty"`
    Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AdAnalyticsGreaterThanOrEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
    return AnalyticsQueryOperator_GTE
}
func (m AdAnalyticsGreaterThanOrEqualFilter) MarshalJSON() ([]byte, error) {
    type M AdAnalyticsGreaterThanOrEqualFilter
    x := struct {
        Operator string `json:"operator"`
        M
    }{M: M(m)}

    x.Operator = "GTE"

    return json.Marshal(x)
}


