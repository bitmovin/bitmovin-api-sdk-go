package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// AnalyticsCountQueryRequest model
type AnalyticsCountQueryRequest struct {
	// Start of timeframe which is queried in UTC format.
	Start *DateTime `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format.
	End *DateTime `json:"end,omitempty"`
	// Analytics license key (required)
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Analytics Query Filters  Each filter requires 3 properties: name, operator and value.   Valid operators are [IN, EQ, NE, LT, LTE, GT, GTE, CONTAINS, NOTCONTAINS]
	Filters   []AnalyticsAbstractFilter `json:"filters,omitempty"`
	OrderBy   []AnalyticsOrderByEntry   `json:"orderBy,omitempty"`
	Dimension AnalyticsAttribute        `json:"dimension,omitempty"`
	Interval  AnalyticsInterval         `json:"interval,omitempty"`
	GroupBy   []AnalyticsAttribute      `json:"groupBy,omitempty"`
	// Whether context data should be included in the response
	IncludeContext *bool `json:"includeContext,omitempty"`
	// Maximum number of rows returned (max. 200)
	Limit *int64 `json:"limit,omitempty"`
	// Offset of data
	Offset *int64 `json:"offset,omitempty"`
}

// UnmarshalJSON unmarshals model AnalyticsCountQueryRequest from a JSON structure
func (m *AnalyticsCountQueryRequest) UnmarshalJSON(raw []byte) error {
	var data struct {
		Start          *DateTime               `json:"start"`
		End            *DateTime               `json:"end"`
		LicenseKey     *string                 `json:"licenseKey"`
		Filters        json.RawMessage         `json:"filters"`
		OrderBy        []AnalyticsOrderByEntry `json:"orderBy"`
		Dimension      AnalyticsAttribute      `json:"dimension"`
		Interval       AnalyticsInterval       `json:"interval"`
		GroupBy        []AnalyticsAttribute    `json:"groupBy"`
		IncludeContext *bool                   `json:"includeContext"`
		Limit          *int64                  `json:"limit"`
		Offset         *int64                  `json:"offset"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result AnalyticsCountQueryRequest

	result.Start = data.Start
	result.End = data.End
	result.LicenseKey = data.LicenseKey
	result.OrderBy = data.OrderBy
	result.Dimension = data.Dimension
	result.Interval = data.Interval
	result.GroupBy = data.GroupBy
	result.IncludeContext = data.IncludeContext
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
