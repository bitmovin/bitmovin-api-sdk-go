package model

// PccVerdict : What a combination says once every session that measured it has been read. Five of the nine answer for the measurement rather than for the device; `aboutTheDevice` says which, and folding those into \"not supported\" is how this data gets misread.
type PccVerdict string

// List of possible PccVerdict values
const (
	PccVerdict_PLAYED                   PccVerdict = "played"
	PccVerdict_CLAIMED_BUT_NOT_VERIFIED PccVerdict = "claimed-but-not-verified"
	PccVerdict_DECLARES_NO_SUPPORT      PccVerdict = "declares-no-support"
	PccVerdict_INCONSISTENT_CLAIM       PccVerdict = "inconsistent-claim"
	PccVerdict_INCONCLUSIVE             PccVerdict = "inconclusive"
	PccVerdict_UNMEASURED               PccVerdict = "unmeasured"
	PccVerdict_NOT_APPLICABLE           PccVerdict = "not-applicable"
	PccVerdict_INFRASTRUCTURE_FAULT     PccVerdict = "infrastructure-fault"
	PccVerdict_NEVER_REACHED            PccVerdict = "never-reached"
)
