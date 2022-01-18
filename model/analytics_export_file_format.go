package model

// AnalyticsExportFileFormat : Analytics Export File Format
type AnalyticsExportFileFormat string

// List of possible AnalyticsExportFileFormat values
const (
	AnalyticsExportFileFormat_CSV     AnalyticsExportFileFormat = "CSV"
	AnalyticsExportFileFormat_PARQUET AnalyticsExportFileFormat = "PARQUET"
)
