package model

// AnalyticsLicenseErrorDetailsConfig model
type AnalyticsLicenseErrorDetailsConfig struct {
	// Are error details enabled on the license
	Enabled *bool `json:"enabled,omitempty"`
	// Is network explorer feature enabled for the license
	NetworkExplorerEnabled *bool `json:"networkExplorerEnabled,omitempty"`
}
