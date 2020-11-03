package model

import (
    "bytes"
    "encoding/json"
    "github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
    "io"
)

// AnalyticsVarianceQueryRequest model
type AnalyticsVarianceQueryRequest struct {
    // Analytics license key (required)
    LicenseKey *string `json:"licenseKey,omitempty"`
    Filters []AnalyticsAbstractFilter `json:"filters,omitempty"`
    OrderBy []AnalyticsOrderByEntry `json:"orderBy,omitempty"`
    Dimension AnalyticsAttribute `json:"dimension,omitempty"`
    Interval AnalyticsInterval `json:"interval,omitempty"`
    GroupBy []AnalyticsAttribute `json:"groupBy,omitempty"`
    // Whether context data should be included in the response
    IncludeContext *bool `json:"includeContext,omitempty"`
    // Maximum number of rows returned (max. 200)
    Limit *int64 `json:"limit,omitempty"`
    // Offset of data
    Offset *int64 `json:"offset,omitempty"`
    // Start of timeframe which is queried in UTC format.
    Start *DateTime `json:"start,omitempty"`
    // End of timeframe which is queried in UTC format.
    End *DateTime `json:"end,omitempty"`
}

// UnmarshalJSON unmarshals model AnalyticsVarianceQueryRequest from a JSON structure
func (m *AnalyticsVarianceQueryRequest) UnmarshalJSON(raw []byte) error {
    var data struct {
        
        LicenseKey *string `json:"licenseKey"`
        Filters json.RawMessage `json:"filters"`
        OrderBy []AnalyticsOrderByEntry `json:"orderBy"`
        Dimension AnalyticsAttribute `json:"dimension"`
        Interval AnalyticsInterval `json:"interval"`
        GroupBy []AnalyticsAttribute `json:"groupBy"`
        IncludeContext *bool `json:"includeContext"`
        Limit *int64 `json:"limit"`
        Offset *int64 `json:"offset"`
        Start *DateTime `json:"start"`
        End *DateTime `json:"end"`
    }

    buf := bytes.NewBuffer(raw)
    dec := json.NewDecoder(buf)
    dec.UseNumber()

    if err := dec.Decode(&data); err != nil {
        return err
    }

    var result AnalyticsVarianceQueryRequest
    
    result.LicenseKey = data.LicenseKey
    result.OrderBy = data.OrderBy
    result.Dimension = data.Dimension
    result.Interval = data.Interval
    result.GroupBy = data.GroupBy
    result.IncludeContext = data.IncludeContext
    result.Limit = data.Limit
    result.Offset = data.Offset
    result.Start = data.Start
    result.End = data.End
    
    allOfFilters, err := UnmarshalAnalyticsAbstractFilterSlice(bytes.NewBuffer(data.Filters), bitutils.JSONConsumer())
    if err != nil && err != io.EOF {
        return err
    }

    result.Filters = allOfFilters

    *m = result

    return nil
}



