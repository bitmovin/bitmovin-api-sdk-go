package model

// PccOverview model
type PccOverview struct {
	// Number of selected cells, one per device pool and codec/protection combination. (required)
	Total *float32 `json:"total,omitempty"`
	// Selected cells with a device-answering verdict, including refusals. (required)
	Answered *float32 `json:"answered,omitempty"`
	// Selected cells that played successfully. (required)
	Played *float32 `json:"played,omitempty"`
}
