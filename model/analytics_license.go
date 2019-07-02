package model

type AnalyticsLicense struct {
	// Id of the resource (required)
	Id string `json:"id,omitempty"`
	// Name of the Analytics License
	Name string `json:"name,omitempty"`
	// License Key (required)
	LicenseKey string `json:"licenseKey,omitempty"`
}

