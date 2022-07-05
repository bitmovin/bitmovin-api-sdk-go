package model

// CdnUsage model
type CdnUsage struct {
	// UTC timestamp which marks the beginning of the time period for which the usage statistics are retrieved.
	From *DateTime `json:"from,omitempty"`
	// UTC timestamp which marks the end of the time period for which the usage statistics are retrieved. The end date is exclusive. For example, if end is 2019-03-29T13:05:00Z, the cost and usage data are retrieved from the start date up to, but not including, 2019-03-29T13:05:00Z.
	To *DateTime `json:"to,omitempty"`
	// Storage usage in GB per month.
	StorageUsage *float64 `json:"storageUsage,omitempty"`
	// Transfer usage in GB.
	TransferUsage *float64 `json:"transferUsage,omitempty"`
}
