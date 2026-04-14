package model

import (
	"bytes"
	"encoding/json"
)

// AnalyticsGreaterThanFilter model
type AnalyticsGreaterThanFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
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

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
