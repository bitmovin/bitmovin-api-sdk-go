package model

import (
	"bytes"
	"encoding/json"
)

// AnalyticsNotEqualFilter model
type AnalyticsNotEqualFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AnalyticsNotEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_NE
}
func (m AnalyticsNotEqualFilter) MarshalJSON() ([]byte, error) {
	type M AnalyticsNotEqualFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "NE"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
