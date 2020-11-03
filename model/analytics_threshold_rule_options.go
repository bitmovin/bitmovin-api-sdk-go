package model


// AnalyticsThresholdRuleOptions model
type AnalyticsThresholdRuleOptions struct {
    // Threshold for triggering alert (required)
    Threshold *float64 `json:"threshold,omitempty"`
    // Persistence of threshold violation in minutes (required)
    Persistence *int32 `json:"persistence,omitempty"`
    // Sample size for rule processing (required)
    SampleSize *int32 `json:"sampleSize,omitempty"`
    // Number of minutes without violation after which incident is considered as recovered (required)
    Recovery *int32 `json:"recovery,omitempty"`
    // Interval for repeating notifications (required)
    RepeatInterval *int32 `json:"repeatInterval,omitempty"`
    // Send notification after incident is resolved
    NotifyOnResolved *bool `json:"notifyOnResolved,omitempty"`
}



