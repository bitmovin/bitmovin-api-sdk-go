package model

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"io"
)

// AnalyticsImpressionsQuery model
type AnalyticsImpressionsQuery struct {
	// Start of timeframe which is queried in UTC format. (required)
	Start *DateTime `json:"start,omitempty"`
	// End of timeframe which is queried in UTC format. (required)
	End *DateTime `json:"end,omitempty"`
	// Analytics license key (required)
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Analytics Query Filters  Each filter requires 3 properties: name, operator and value.  Valid operators are [IN, EQ, NE, LT, LTE, GT, GTE, CONTAINS, NOTCONTAINS]
	Filters []AnalyticsAbstractFilter `json:"filters,omitempty"`
	// Number of returned impressions
	Limit *int32 `json:"limit,omitempty"`
}

// UnmarshalJSON unmarshals model AnalyticsImpressionsQuery from a JSON structure
func (m *AnalyticsImpressionsQuery) UnmarshalJSON(raw []byte) error {
	var data struct {
		Start      *DateTime       `json:"start"`
		End        *DateTime       `json:"end"`
		LicenseKey *string         `json:"licenseKey"`
		Filters    json.RawMessage `json:"filters"`
		Limit      *int32          `json:"limit"`
	}

	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var result AnalyticsImpressionsQuery

	result.Start = data.Start
	result.End = data.End
	result.LicenseKey = data.LicenseKey
	result.Limit = data.Limit

	allOfFilters, err := UnmarshalAnalyticsAbstractFilterSlice(bytes.NewBuffer(data.Filters), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	result.Filters = allOfFilters

	*m = result

	return nil
}
