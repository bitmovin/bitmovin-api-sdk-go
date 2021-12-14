package pagination

import (
	"bytes"
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
	"io"
)

// AdaptationSetsListPagination model
type AdaptationSetsListPagination struct {
	TotalCount int64                 `json:"totalCount,omitempty"`
	Offset     int32                 `json:"offset,omitempty"`
	Limit      int32                 `json:"limit,omitempty"`
	Previous   string                `json:"previous,omitempty"`
	Next       string                `json:"next,omitempty"`
	Items      []model.AdaptationSet `json:"items,omitempty"`
}

// UnmarshalJSON unmarshals pagination model AdaptationSetsListPagination from a JSON structure
func (m *AdaptationSetsListPagination) UnmarshalJSON(b []byte) error {
	var pageResp model.PaginationResponse
	if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}
	items, err := model.UnmarshalAdaptationSetSlice(bytes.NewBuffer(pageResp.Items), bitutils.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}
	var result AdaptationSetsListPagination

	result.TotalCount = bitutils.Int64Value(pageResp.TotalCount)
	result.Offset = bitutils.Int32Value(pageResp.Offset)
	result.Limit = bitutils.Int32Value(pageResp.Limit)
	result.Previous = bitutils.StringValue(pageResp.Previous)
	result.Next = bitutils.StringValue(pageResp.Next)
	result.Items = items

	*m = result

	return nil
}
