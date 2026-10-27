package model

// LiveDashManifest model
type LiveDashManifest struct {
	// Dash manifest id (required)
	ManifestId *string `json:"manifestId,omitempty"`
	// Timeshift in seconds
	Timeshift *float64 `json:"timeshift,omitempty"`
	// Live edge offset in seconds. How it is applied depends on the type of the segmented representations in the manifest. For `TIMELINE` representations, new segments are held back for this duration (rounded up to whole segments) before they are added to the SegmentTimeline, and are always added in order. This adds the same amount of latency, but a segment that arrives late by up to this duration does not interrupt the manifest. A segment that arrives later than that stops the SegmentTimeline from updating until it arrives. If it does not arrive at all, a placeholder segment is added in its place once a few newer segments are available. If set to 0, segments are added as soon as they are available, still in order. For `TEMPLATE` representations, the `availabilityStartTime` is set this many seconds later, so players request each segment this much later. Late segments do not change the manifest; a player that requests a segment before it is uploaded gets an error. Segments are also kept in the manifest for this duration on top of the timeshift. We recommend a value of at least 2 segment durations.
	LiveEdgeOffset *float64 `json:"liveEdgeOffset,omitempty"`
	// The suggestedPresentationDelay to be set in the DASH manifest. If nothing is set, no value will be set.
	SuggestedPresentationDelay *float64 `json:"suggestedPresentationDelay,omitempty"`
	// The minimumUpdatePeriod to be set in the DASH manifest. If nothing is set, the segment duration will be set.
	MinimumUpdatePeriod *float64 `json:"minimumUpdatePeriod,omitempty"`
	// The mode to trigger the availabilityStartTime initialization.
	AvailabilityStartTimeMode AvailabilityStartTimeMode     `json:"availabilityStartTimeMode,omitempty"`
	AdMarkerSettings          *DashManifestAdMarkerSettings `json:"adMarkerSettings,omitempty"`
}
