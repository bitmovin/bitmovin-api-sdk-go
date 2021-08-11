package model

// AnalyticsErrorData model
type AnalyticsErrorData struct {
	// Exception message
	ExceptionMessage *string `json:"exceptionMessage,omitempty"`
	// Additional error data
	AdditionalData      *string  `json:"additionalData,omitempty"`
	ExceptionStacktrace []string `json:"exceptionStacktrace,omitempty"`
}
