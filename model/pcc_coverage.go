package model

// PccCoverage model
type PccCoverage struct {
	// Distinct pools in the held measurement before any view exclusions. (required)
	DevicesBeforeView *int32 `json:"devicesBeforeView,omitempty"`
	// Pools omitted by the device filter after prerelease exclusions. Each omitted pool is counted once, in prerelease, device-filter, then reported-only order. (required)
	DevicesExcludedByDeviceFilter *int32 `json:"devicesExcludedByDeviceFilter,omitempty"`
	// Pools omitted by reportedOnly because no selected cell answers about the device, after prerelease and device-filter exclusions. (required)
	DevicesExcludedAsUnreported *int32 `json:"devicesExcludedAsUnreported,omitempty"`
	// Distinct pools omitted because all recorded browser evidence is pre-release. Zero when includePrerelease is true. Unknown browser identities do not cause prerelease exclusion. Row coverage and summaries describe the retained pools. (required)
	DevicesExcludedAsPrerelease *int32   `json:"devicesExcludedAsPrerelease,omitempty"`
	Devices                     *float32 `json:"devices,omitempty"`
	// Pools with no included sessions. Present in the report as unmeasured, including when the start date excluded all evidence. (required)
	DevicesWithNoSession *float32 `json:"devicesWithNoSession,omitempty"`
	// Pools every one of whose sessions came from one physical machine — a claim about that machine, not the model. (required)
	DevicesOnOneUnit *float32 `json:"devicesOnOneUnit,omitempty"`
	// Pools no naming rule recognised, published under the stated placeholder. Counted here so a reader can tell how much of the fleet this report cannot name rather than discovering it row by row. (required)
	DevicesUnderPlaceholderName *float32 `json:"devicesUnderPlaceholderName,omitempty"`
	// Cells whose verdict rests on a single session. (required)
	CellsOnOneSession *float32 `json:"cellsOnOneSession,omitempty"`
	// Jobs the fleet could not attribute to any pool, so no row of this report accounts for them. (required)
	UnattributableJobs *float32 `json:"unattributableJobs,omitempty"`
	// Jobs that had not finished when this was assembled. Their pools carry nothing measured. (required)
	UnsettledJobs *float32 `json:"unsettledJobs,omitempty"`
	// Pools with sessions excluded by the start date or session limit, including those with no included evidence. (required)
	DevicesOnRecentEvidence *float32 `json:"devicesOnRecentEvidence,omitempty"`
}
