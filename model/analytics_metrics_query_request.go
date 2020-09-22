package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// AnalyticsMetricsQueryRequest model
type AnalyticsMetricsQueryRequest struct {
	// Start of timeframe which is queried in UTC format.
	Start *DateTime `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format.
	End *DateTime `json:"end,omitempty"`
	// Analytics license key (required)
	LicenseKey *string                   `json:"licenseKey,omitempty"`
	Filters    []AnalyticsAbstractFilter `json:"filters,omitempty"`
	OrderBy    []AnalyticsOrderByEntry   `json:"orderBy,omitempty"`
	Interval   AnalyticsInterval         `json:"interval,omitempty"`
	GroupBy    []AnalyticsAttribute      `json:"groupBy,omitempty"`
	// Maximum number of rows returned (max. 200)
	Limit *int64 `json:"limit,omitempty"`
	// Offset of data used for pagination
	Offset *int64 `json:"offset,omitempty"`
}

// UnmarshalJSON unmarshals model AnalyticsMetricsQueryRequest from a JSON structure
func (m *AnalyticsMetricsQueryRequest) UnmarshalJSON(raw []byte) error {
	var data struct {
		Start      *DateTime               `json:"start"`
		End        *DateTime               `json:"end"`
		LicenseKey *string                 `json:"licenseKey"`
		Filters    json.RawMessage         `json:"filters"`
		OrderBy    []AnalyticsOrderByEntry `json:"orderBy"`
		Interval   AnalyticsInterval       `json:"interval"`
		GroupBy    []AnalyticsAttribute    `json:"groupBy"`
		Limit      *int64                  `json:"limit"`
		Offset     *int64                  `json:"offset"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result AnalyticsMetricsQueryRequest

	result.Start = data.Start
	result.End = data.End
	result.LicenseKey = data.LicenseKey
	result.OrderBy = data.OrderBy
	result.Interval = data.Interval
	result.GroupBy = data.GroupBy
	result.Limit = data.Limit
	result.Offset = data.Offset

	allOfFilters, err := UnmarshalAnalyticsAbstractFilterSlice(bytes.NewBuffer(data.Filters), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Filters = allOfFilters

	*m = result

	return nil
}
