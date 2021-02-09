package model

// AnalyticsOutputType : AnalyticsOutputType model
type AnalyticsOutputType string

// List of possible AnalyticsOutputType values
const (
	AnalyticsOutputType_S3_ROLE_BASED       AnalyticsOutputType = "S3_ROLE_BASED"
	AnalyticsOutputType_GCS_SERVICE_ACCOUNT AnalyticsOutputType = "GCS_SERVICE_ACCOUNT"
)
