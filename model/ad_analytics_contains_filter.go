package model

import (
	"bytes"
	"encoding/json"
)

// AdAnalyticsContainsFilter model
type AdAnalyticsContainsFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AdAnalyticsContainsFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_CONTAINS
}
func (m AdAnalyticsContainsFilter) MarshalJSON() ([]byte, error) {
	type M AdAnalyticsContainsFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "CONTAINS"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
