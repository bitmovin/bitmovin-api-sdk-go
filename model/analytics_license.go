package model

// AnalyticsLicense model
type AnalyticsLicense struct {
	// Id of the Analytics License
	Id *string `json:"id,omitempty"`
	// Creation date of the Analytics License, returned as ISO 8601 date-time format
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// User-specific meta data. This can hold anything.
	CustomData *map[string]interface{} `json:"customData,omitempty"`
	// License Key
	LicenseKey *string `json:"licenseKey,omitempty"`
	// Name of the Analytics License
	Name *string `json:"name,omitempty"`
	// The industry of the organization associated with the Analytics License
	Industry *string `json:"industry,omitempty"`
	// The subindustry of the organization associated with the Analytics License
	SubIndustry *string `json:"subIndustry,omitempty"`
	// Whether the Do Not Track request from the browser should be ignored
	IgnoreDNT *bool `json:"ignoreDNT,omitempty"`
	// Number of impressions recorded
	Impressions *int64 `json:"impressions,omitempty"`
	// Maximum number of impressions
	MaxImpressions *int64 `json:"maxImpressions,omitempty"`
	// The timezone of the Analytics License
	TimeZone *string `json:"timeZone,omitempty"`
	// Retention time of impressions, returned as ISO 8601 duration format: P(n)Y(n)M(n)DT(n)H(n)M(n)S
	RetentionTime *string `json:"retentionTime,omitempty"`
	// Retention time for compressed data, returned as ISO 8601 duration format: P(n)Y(n)M(n)DT(n)H(n)M(n)S
	CompressedRetentionTime *string `json:"compressedRetentionTime,omitempty"`
	// Whitelisted domains
	Domains []AnalyticsLicenseDomain `json:"domains,omitempty"`
	// Whether the data of this license should be included in the industry insights or not
	IncludeInInsights *bool `json:"includeInInsights,omitempty"`
	// Labels for CustomData fields
	CustomDataFieldLabels *AnalyticsLicenseCustomDataFieldLabels `json:"customDataFieldLabels,omitempty"`
	// The number of customData fields available
	CustomDataFieldsCount *int32 `json:"customDataFieldsCount,omitempty"`
	// Order index of license
	OrderIndex *int32 `json:"orderIndex,omitempty"`
	// The rate limit of this license
	RateLimit *string                   `json:"rateLimit,omitempty"`
	Features  *AnalyticsLicenseFeatures `json:"features,omitempty"`
	// The expiration date of the license if applicable, returned as ISO 8601 date-time format
	PlanExpiredAt *DateTime `json:"planExpiredAt,omitempty"`
}
