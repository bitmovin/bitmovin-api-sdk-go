package model

// AnalyticsErrorDetail model
type AnalyticsErrorDetail struct {
	// Error id
	ErrorId *int64 `json:"errorId,omitempty"`
	// Server timestamp
	Time *DateTime `json:"time,omitempty"`
	// Client timestamp
	ClientTime *DateTime `json:"clientTime,omitempty"`
	// Error code
	Code *int32 `json:"code,omitempty"`
	// Error message
	Message   *string             `json:"message,omitempty"`
	ErrorData *AnalyticsErrorData `json:"errorData,omitempty"`
	// Severity of the error
	Severity     *string                `json:"severity,omitempty"`
	HttpRequests []AnalyticsHttpRequest `json:"httpRequests,omitempty"`
}
