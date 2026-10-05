package model

// ResetLiveManifestTimeShift model
type ResetLiveManifestTimeShift struct {
	// Id of the resource (required)
	Id *string `json:"id,omitempty"`
	// Specifies how many seconds of content remain in the manifest after older segments are removed. At least one segment is always retained. If neither `residualPeriodInSeconds` nor `offsetInSeconds` is set, all segments except the most recent are removed. For DASH manifests that use SegmentTemplate, the retained duration also includes the value configured for `liveEdgeOffset`.
	ResidualPeriodInSeconds *float64 `json:"residualPeriodInSeconds,omitempty"`
	// Specifies an offset, in seconds, from the start of the live event. All segments before this position are removed from the affected manifests. For example, assume a segment length of 2 seconds and a configured `timeshift` of 120 seconds (2 minutes). If the most recent segment is `segment_80.ts`, the manifest contains 60 segments, from `segment_21.ts` through `segment_80.ts`. Setting `offsetInSeconds` to `120` sets the target segment number to 60 (`targetSegmentNumber = offsetInSeconds / segmentLength`). All segments before `segment_60.ts` are removed. Each affected manifest then contains `segment_60.ts` through `segment_80.ts`.  *Note:* Do not set both `offsetInSeconds` and `residualPeriodInSeconds`.
	OffsetInSeconds *float64 `json:"offsetInSeconds,omitempty"`
	// The IDs of the manifests to update. If omitted, all supported manifests associated with the encoding are updated. HLS live manifests are supported. DASH live manifests require encoder version 2.235.0 or later.
	ManifestIds []string `json:"manifestIds,omitempty"`
	// If set to true, the Progressive muxing start position will be shifted to the start of the first remaining segment after the removal.  NOTE: This only works for Progressive MP4 muxings.
	ShiftProgressiveMuxingStartPosition *bool `json:"shiftProgressiveMuxingStartPosition,omitempty"`
}
