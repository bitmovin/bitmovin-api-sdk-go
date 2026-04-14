package model

import (
	"bytes"
	"encoding/json"
)

// AdAnalyticsLessThanFilter model
type AdAnalyticsLessThanFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AdAnalyticsLessThanFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_LT
}
func (m AdAnalyticsLessThanFilter) MarshalJSON() ([]byte, error) {
	type M AdAnalyticsLessThanFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "LT"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
