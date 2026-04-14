package model

import (
	"bytes"
	"encoding/json"
)

// AnalyticsGreaterThanOrEqualFilter model
type AnalyticsGreaterThanOrEqualFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsGreaterThanOrEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_GTE
}
func (m AnalyticsGreaterThanOrEqualFilter) MarshalJSON() ([]byte, error) {
	type M AnalyticsGreaterThanOrEqualFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "GTE"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
