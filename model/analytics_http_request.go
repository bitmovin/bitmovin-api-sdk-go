package model

// AnalyticsHttpRequest model
type AnalyticsHttpRequest struct {
	// Client timestamp
	ClientTime *DateTime                `json:"clientTime,omitempty"`
	Type       AnalyticsHttpRequestType `json:"type,omitempty"`
	// Http request URL
	Url *string `json:"url,omitempty"`
	// Last redirect location
	LastRedirectLocation *string `json:"lastRedirectLocation,omitempty"`
	// Http request status
	HttpStatus *int32 `json:"httpStatus,omitempty"`
	// Download time
	DownloadTime *int64 `json:"downloadTime,omitempty"`
	// Time to first byte
	TimeToFirstByte *int64 `json:"timeToFirstByte,omitempty"`
	// Size
	Size *int64 `json:"size,omitempty"`
	// Was http request successful
	Success *bool `json:"success,omitempty"`
}
