package model

type TimeSpan struct {
	// Start offset of the time frame in milliseconds
	From *int64 `json:"from,omitempty"`
	// End offset of the time frame in milliseconds
	To *int64 `json:"to,omitempty"`
}

