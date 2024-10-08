package model

// ResetLiveManifestTimeShift model
type ResetLiveManifestTimeShift struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Determines how many seconds will be left in the manifest after segments are removed. If this is not set, all but one segment will be removed.
	ResidualPeriodInSeconds *float64 `json:"residualPeriodInSeconds,omitempty"`
	// Offset in seconds from the start of the live event, defining the point from which all segments within that duration will be removed from the given manifests. E.g.: The segment length is 2 seconds and a timeshift of 120 seconds (2 minutes) is configured.  The manifest contains 60 segments with the last segment number being 80 (`segment_80.ts`).  This means the manifest contains `segment_20.ts` to `segment_80.ts` (timeshift of 2 minutes equals 60 segments in manifest) If you set `offsetInSeconds` to `120`, all segments below segment number 60 (`segment_60.ts`) will be removed. (`targetSegmentNumber = offsetInSeconds / segmentLength`) The manifests will then contain `segment_60.ts` to `segment_80.ts` *Note:* Only `offsetInSeconds` or `residualPeriodInSeconds` can be set.
	OffsetInSeconds *float64 `json:"offsetInSeconds,omitempty"`
	// The ids of the manifests to update. If this property is not set, all manifests tied to the encoding are updated.
	ManifestIds []string `json:"manifestIds,omitempty"`
	// If set to true, the Progressive muxing start position will be shifted to the start of the first remaining segment after the removal.  NOTE: This only works for Progressive MP4 muxings.
	ShiftProgressiveMuxingStartPosition *bool `json:"shiftProgressiveMuxingStartPosition,omitempty"`
}
