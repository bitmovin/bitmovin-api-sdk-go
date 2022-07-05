package model

// CdnUsageStatistics model
type CdnUsageStatistics struct {
	// UTC timestamp which marks the beginning of the time period for which the usage statistics are retrieved.
	From *DateTime `json:"from,omitempty"`
	// UTC timestamp which marks the end of the time period for which the usage statistics are retrieved. The end date is exclusive. For example, if end is 2019-03-28T13:05:00Z, the cost and usage data are retrieved from the start date up to, but not including, 2019-03-28T13:05:00Z.
	To *DateTime `json:"to,omitempty"`
	// Total storage usage in GB per month.
	StorageUsageTotal *float64 `json:"storageUsageTotal,omitempty"`
	// Total transfer usage in GB.
	TransferUsageTotal *float64   `json:"transferUsageTotal,omitempty"`
	Usage              []CdnUsage `json:"usage,omitempty"`
}
