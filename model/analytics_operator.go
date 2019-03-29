package model
type AnalyticsOperator string

// List of AnalyticsOperator
const (
	AnalyticsOperator_EQ AnalyticsOperator = "EQ"
	AnalyticsOperator_NE AnalyticsOperator = "NE"
	AnalyticsOperator_LT AnalyticsOperator = "LT"
	AnalyticsOperator_LTE AnalyticsOperator = "LTE"
	AnalyticsOperator_GT AnalyticsOperator = "GT"
	AnalyticsOperator_GTE AnalyticsOperator = "GTE"
	AnalyticsOperator_CONTAINS AnalyticsOperator = "CONTAINS"
	AnalyticsOperator_NOTCONTAINS AnalyticsOperator = "NOTCONTAINS"
)