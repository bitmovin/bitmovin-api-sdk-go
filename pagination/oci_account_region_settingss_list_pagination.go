package pagination

import (
	"encoding/json"
	"github.com/bitmovin/bitmovin-api-sdk-go/bitutils"
	"github.com/bitmovin/bitmovin-api-sdk-go/model"
)

// OciAccountRegionSettingssListPagination model
type OciAccountRegionSettingssListPagination struct {
	TotalCount int64                            `json:"totalCount,omitempty"`
	Offset     int32                            `json:"offset,omitempty"`
	Limit      int32                            `json:"limit,omitempty"`
	Previous   string                           `json:"previous,omitempty"`
	Next       string                           `json:"next,omitempty"`
	Items      []model.OciAccountRegionSettings `json:"items,omitempty"`
}

// UnmarshalJSON unmarshals pagination model OciAccountRegionSettingssListPagination from a JSON structure
func (m *OciAccountRegionSettingssListPagination) UnmarshalJSON(b []byte) error {
	var pageResp model.PaginationResponse
	if err := json.Unmarshal(b, &pageResp); err != nil {
		return err
	}

	var items []model.OciAccountRegionSettings
	if err := json.Unmarshal(pageResp.Items, &items); err != nil {
		return err
	}
	var result OciAccountRegionSettingssListPagination

	result.TotalCount = bitutils.Int64Value(pageResp.TotalCount)
	result.Offset = bitutils.Int32Value(pageResp.Offset)
	result.Limit = bitutils.Int32Value(pageResp.Limit)
	result.Previous = bitutils.StringValue(pageResp.Previous)
	result.Next = bitutils.StringValue(pageResp.Next)
	result.Items = items

	*m = result

	return nil
}
