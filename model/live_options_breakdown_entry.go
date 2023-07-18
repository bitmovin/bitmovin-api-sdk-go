package model

// LiveOptionsBreakdownEntry model
type LiveOptionsBreakdownEntry struct {
	// Date, format: yyyy-MM-dd (required)
	Date *Date             `json:"date,omitempty"`
	Hd   *LiveOptionsEntry `json:"hd,omitempty"`
}
