package model

// AdvisoryConfidence : AdvisoryConfidence model
type AdvisoryConfidence string

// List of possible AdvisoryConfidence values
const (
	AdvisoryConfidence_HIGH    AdvisoryConfidence = "HIGH"
	AdvisoryConfidence_MEDIUM  AdvisoryConfidence = "MEDIUM"
	AdvisoryConfidence_LOW     AdvisoryConfidence = "LOW"
	AdvisoryConfidence_UNKNOWN AdvisoryConfidence = "UNKNOWN"
)
