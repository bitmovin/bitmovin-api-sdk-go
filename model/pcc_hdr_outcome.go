package model

// PccHdrOutcome : PccHdrOutcome model
type PccHdrOutcome string

// List of possible PccHdrOutcome values
const (
	PccHdrOutcome_HDR           PccHdrOutcome = "hdr"
	PccHdrOutcome_SDR           PccHdrOutcome = "sdr"
	PccHdrOutcome_CLAIMED       PccHdrOutcome = "claimed"
	PccHdrOutcome_DENIED        PccHdrOutcome = "denied"
	PccHdrOutcome_UNESTABLISHED PccHdrOutcome = "unestablished"
	PccHdrOutcome_UNREPORTED    PccHdrOutcome = "unreported"
)
