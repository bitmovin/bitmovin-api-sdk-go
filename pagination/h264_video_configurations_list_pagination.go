package pagination

import (
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// H264VideoConfigurationsListPagination model
type H264VideoConfigurationsListPagination struct {
	TotalCount int64                          `json:"totalCount,omitempty"`
	Offset     int32                          `json:"offset,omitempty"`
	Limit      int32                          `json:"limit,omitempty"`
	Previous   string                         `json:"previous,omitempty"`
	Next       string                         `json:"next,omitempty"`
	Items      []model.H264VideoConfiguration `json:"items,omitempty"`
}

// UnmarshalJSON unmarshals pagination model H264VideoConfigurationsListPagination from a JSON structure
func (m *H264VideoConfigurationsListPagination) UnmarshalJSON(b []byte) error {
	var pageResp model.PaginationResponse
	if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

	var items []model.H264VideoConfiguration
	if err := json.Unmarshal(pageResp.Items, &items); err != nil {
		return err
	}
	var result H264VideoConfigurationsListPagination

	result.TotalCount = bitutils.Int64Value(pageResp.TotalCount)
	result.Offset = bitutils.Int32Value(pageResp.Offset)
	result.Limit = bitutils.Int32Value(pageResp.Limit)
	result.Previous = bitutils.StringValue(pageResp.Previous)
	result.Next = bitutils.StringValue(pageResp.Next)
	result.Items = items

	*m = result

	return nil
}
