package model

// A key event in the analyzed media, supported by one or more active observations
type AiSceneAnalysisLiveKeyMoment struct {
	// Opaque key moment ID unique within the Live Analysis. Stable through refinements and never reused after withdrawal. (required)
	Id *string `json:"id,omitempty"`
	// Result sequence that introduced or last changed this moment's importance level, times or observation references. Unchanged moments retain this value. (required)
	LastUpdatedSequence *int64 `json:"lastUpdatedSequence,omitempty"`
	// Importance level, not a position within the result snapshot: 1 is highest and larger values indicate lower importance. Ties are allowed. Adding another moment does not by itself renumber existing importance levels. There is no upper bound or implied limit on the number of key moments. (required)
	ImportanceLevel *int32 `json:"importanceLevel,omitempty"`
	// Start of the key moment in seconds (required)
	StartInSeconds *float64 `json:"startInSeconds,omitempty"`
	// End of the key moment in seconds (required)
	EndInSeconds *float64 `json:"endInSeconds,omitempty"`
	// Time of the key moment's actual high point in seconds (required)
	AnchorInSeconds *float64 `json:"anchorInSeconds,omitempty"`
	// IDs of active observations in the enclosing result that support this moment. Multiple key moments may reference the same observation. (required)
	Observations []string `json:"observations,omitempty"`
}
