package model


// LiveDashManifest model
type LiveDashManifest struct {
    // Dash manifest ids (required)
    ManifestId *string `json:"manifestId,omitempty"`
    // Timeshift in seconds
    Timeshift *float64 `json:"timeshift,omitempty"`
    // Live edge offset in seconds
    LiveEdgeOffset *float64 `json:"liveEdgeOffset,omitempty"`
    // The suggestedPresentationDelay to be set in the DASH manifest. If nothing is set, no value will be set.
    SuggestedPresentationDelay *float64 `json:"suggestedPresentationDelay,omitempty"`
    // The minimumUpdatePeriod to be set in the DASH manifest. If nothing is set, the segment duration will be set.
    MinimumUpdatePeriod *float64 `json:"minimumUpdatePeriod,omitempty"`
    // The mode to trigger the availabilityStartTime initialization.
    AvailabilityStartTimeMode AvailabilityStartTimeMode `json:"availabilityStartTimeMode,omitempty"`
}



