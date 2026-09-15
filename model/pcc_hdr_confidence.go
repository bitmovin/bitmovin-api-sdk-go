package model

// PccHdrConfidence : Which instrument established the picture, where anything did.
type PccHdrConfidence string

// List of possible PccHdrConfidence values
const (
	PccHdrConfidence_EVIDENCE      PccHdrConfidence = "evidence"
	PccHdrConfidence_CLAIM         PccHdrConfidence = "claim"
	PccHdrConfidence_UNESTABLISHED PccHdrConfidence = "unestablished"
)
