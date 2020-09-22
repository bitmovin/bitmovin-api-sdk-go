package model

import "encoding/json"

type PaginationResponse struct {
	TotalCount *int64          `json:"totalCount,omitempty"`
	Offset     *int32          `json:"offset,omitempty"`
	Limit      *int32          `json:"limit,omitempty"`
	Previous   *string         `json:"previous,omitempty"`
	Next       *string         `json:"next,omitempty"`
	Items      json.RawMessage `json:"items,omitempty"`
}
