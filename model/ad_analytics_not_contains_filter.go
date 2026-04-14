package model

import (
	"bytes"
	"encoding/json"
)

// AdAnalyticsNotContainsFilter model
type AdAnalyticsNotContainsFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
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

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
