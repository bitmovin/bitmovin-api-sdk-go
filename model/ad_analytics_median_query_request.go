package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// AdAnalyticsMedianQueryRequest model
type AdAnalyticsMedianQueryRequest struct {
	// Analytics license key (required)
	LicenseKey *string                     `json:"licenseKey,omitempty"`
	Filters    []AdAnalyticsAbstractFilter `json:"filters,omitempty"`
	OrderBy    []AdAnalyticsOrderByEntry   `json:"orderBy,omitempty"`
	Dimension  AdAnalyticsAttribute        `json:"dimension,omitempty"`
	Interval   AnalyticsInterval           `json:"interval,omitempty"`
	GroupBy    []AdAnalyticsAttribute      `json:"groupBy,omitempty"`
	// Maximum number of rows returned (max. 200)
	Limit *int64 `json:"limit,omitempty"`
	// Offset of data
	Offset *int64 `json:"offset,omitempty"`
	// Start of timeframe which is queried in UTC format.
	Start *DateTime `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format.
	End *DateTime `json:"end,omitempty"`
}

// UnmarshalJSON unmarshals model AdAnalyticsMedianQueryRequest from a JSON structure
func (m *AdAnalyticsMedianQueryRequest) UnmarshalJSON(raw []byte) error {
	var data struct {
		LicenseKey *string                   `json:"licenseKey"`
		Filters    json.RawMessage           `json:"filters"`
		OrderBy    []AdAnalyticsOrderByEntry `json:"orderBy"`
		Dimension  AdAnalyticsAttribute      `json:"dimension"`
		Interval   AnalyticsInterval         `json:"interval"`
		GroupBy    []AdAnalyticsAttribute    `json:"groupBy"`
		Limit      *int64                    `json:"limit"`
		Offset     *int64                    `json:"offset"`
		Start      *DateTime                 `json:"start"`
		End        *DateTime                 `json:"end"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result AdAnalyticsMedianQueryRequest

	result.LicenseKey = data.LicenseKey
	result.OrderBy = data.OrderBy
	result.Dimension = data.Dimension
	result.Interval = data.Interval
	result.GroupBy = data.GroupBy
	result.Limit = data.Limit
	result.Offset = data.Offset
	result.Start = data.Start
	result.End = data.End

	allOfFilters, err := UnmarshalAdAnalyticsAbstractFilterSlice(bytes.NewBuffer(data.Filters), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Filters = allOfFilters

	*m = result

	return nil
}
