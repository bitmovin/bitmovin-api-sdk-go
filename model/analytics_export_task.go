package model

type AnalyticsExportTask struct {
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Start of timeframe which is exported in UTC format
	StartTime string `json:"startTime,omitempty"`
	// End of timeframe which is exported in UTC format
	EndTime string `json:"endTime,omitempty"`
	// Name of the export task
	Name string `json:"name,omitempty"`
	// Export task description
	Description string `json:"description,omitempty"`
	// License key
	LicenseKey string `json:"licenseKey,omitempty"`
	Output *AnalyticsExportTaskOutputTarget `json:"output,omitempty"`
}

