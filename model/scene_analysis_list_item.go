package model

// SceneAnalysisListItem model
type SceneAnalysisListItem struct {
	// AI scene analysis ID (required)
	Id *string `json:"id,omitempty"`
	// ID of the associated encoding (required)
	EncodingId *string `json:"encodingId,omitempty"`
	// Creation timestamp, returned as UTC in ISO 8601 format: YYYY-MM-DDThh:mm:ssZ (required)
	CreatedAt *DateTime `json:"createdAt,omitempty"`
	// Analysis description. Empty when analysis metadata is unavailable (required)
	Description *string `json:"description,omitempty"`
	// Inferred title representing the analyzed content as a whole. If omitted or null, the title is not available.
	Title *string `json:"title,omitempty"`
	// Analysis keywords in their original order and casing, including duplicates. Omitted or empty when analysis metadata is unavailable; consumers must treat both representations as an empty list
	Keywords []string `json:"keywords,omitempty"`
	// Number of scenes in the analysis. Zero when analysis metadata is unavailable (required)
	SceneCount *int32 `json:"sceneCount,omitempty"`
	// Unique language codes for available translated analysis details in backend-defined deterministic order. Order and casing are returned unchanged. Omitted or empty when no translations are available; consumers must treat both representations as an empty list
	OutputLanguageCodes []string `json:"outputLanguageCodes,omitempty"`
	// The scene segment that best matches searchText. Present only for semantic-search requests with a non-blank searchText; omitted from ordinary list results.
	MatchingSegment *SceneAnalysisMatchingSegment `json:"matchingSegment,omitempty"`
}
