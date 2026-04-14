package model

import (
	"bytes"
	"encoding/json"
)

// AnalyticsLessThanOrEqualFilter model
type AnalyticsLessThanOrEqualFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
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

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
