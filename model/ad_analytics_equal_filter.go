package model

import (
	"bytes"
	"encoding/json"
)

// AdAnalyticsEqualFilter model
type AdAnalyticsEqualFilter struct {
	Name  *map[string]interface{} `json:"name,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

func (m AdAnalyticsEqualFilter) AnalyticsQueryOperator() AnalyticsQueryOperator {
	return AnalyticsQueryOperator_EQ
}
func (m AdAnalyticsEqualFilter) MarshalJSON() ([]byte, error) {
	type M AdAnalyticsEqualFilter
	x := struct {
		Operator string `json:"operator"`
		M
	}{M: M(m)}

	x.Operator = "EQ"

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(x); err != nil {
		return nil, err
	}
	return bytes.TrimRight(buf.Bytes(), "\n"), nil
}
