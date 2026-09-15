package model

// PccCompatibilityMatrix model
type PccCompatibilityMatrix struct {
	// The effective selection already applied to this matrix, its summary and coverage. (required)
	View *PccReportView `json:"view,omitempty"`
	// When this report was assembled from what had been measured by then. (required)
	AssembledAt *string `json:"assembledAt,omitempty"`
	// Every Bitmovin Player version that measured any device here. More than one means the measurement spanned a player release, and support is a property of the player and the device together. (required)
	PlayerVersions []string `json:"playerVersions,omitempty"`
	// UUIDs of matching runs whose job metadata was read, including runs with no included session evidence. Never run names. Quote one to Bitmovin support while the fleet still holds it. The count says nothing about coverage. (required)
	RunIds []string `json:"runIds,omitempty"`
	// Inclusive run creation instant in UTC used by this generation, or null for no cutoff. Legacy dates mean midnight UTC. Changing the held cutoff does not alter this report. (required)
	StartDate *DateTime `json:"startDate,omitempty"`
	// Maximum sessions read per pool by this generation. Null for reports produced before a limit was recorded. (required)
	SessionLimit *int32       `json:"sessionLimit,omitempty"`
	Coverage     *PccCoverage `json:"coverage,omitempty"`
	// Every verdict mark and its wording, so the grid reads without this service's source. (required)
	Legend []PccVerdictLegendEntry `json:"legend,omitempty"`
	// The same for the marks an HDR result wears. (required)
	HdrLegend    []PccHdrLegendEntry `json:"hdrLegend,omitempty"`
	Combinations []PccCombination    `json:"combinations,omitempty"`
	Devices      []PccDevice         `json:"devices,omitempty"`
	Summary      *PccSummary         `json:"summary,omitempty"`
}
