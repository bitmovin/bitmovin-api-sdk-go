package model

type AnalyticsExportTaskDetails struct {
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
	// Id of the resource
	Id string `json:"id,omitempty"`
	// Progress of the export task
	Progress *int32 `json:"progress,omitempty"`
	Status AnalyticsExportStatus `json:"status,omitempty"`
	// UTC timestamp when the export task started
	StartedAt string `json:"startedAt,omitempty"`
	// UTC timestamp when the export task finished
	FinishedAt string `json:"finishedAt,omitempty"`
}

