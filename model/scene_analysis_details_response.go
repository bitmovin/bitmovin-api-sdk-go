package model

// VOD analysis of an encoded video. Field availability depends on enabled features and the analysis version. The REST response includes only fields defined by this API; asset-description.json files in customer outputs can contain additional fields. Empty arrays are omitted. Some unavailable values are returned as null, as noted on individual fields. Empty or missing values alone do not establish whether a feature was assessed. In language-specific results, free-text fields are translated into the requested language unless otherwise noted.
type SceneAnalysisDetailsResponse struct {
	// Scenes identified in the video. (required)
	Scenes []Scene `json:"scenes,omitempty"`
	// Duration of the analyzed video in seconds.
	Duration *float64 `json:"duration,omitempty"`
	// AI-generated summary of the video as a whole. (required)
	Description *string `json:"description,omitempty"`
	// Inferred title representing the analyzed content as a whole. If omitted or null, the title is not available.
	Title *string `json:"title,omitempty"`
	// Free-text keywords describing the video as a whole. Preserve the returned spelling and case. (required)
	Keywords []string `json:"keywords,omitempty"`
	// AI-estimated content ratings for the indicated regions and rating systems. (required)
	Ratings []Rating `json:"ratings,omitempty"`
	// Sensitive topics for the video as a whole, in free text. This list can differ from the combined scene-level topics. (required)
	SensitiveTopics []string `json:"sensitiveTopics,omitempty"`
	// Distinct IAB Content Taxonomy identifiers drawn from the scenes' IAB sensitive-topic annotations. The list can be empty or unavailable even when individual scenes have topic annotations. (required)
	IabSensitiveTopicTaxonomies []string `json:"iabSensitiveTopicTaxonomies,omitempty"`
	// Language codes identified for spoken dialogue or narration in the source content. Theme-song vocals, lyrics, music and written text are excluded. The list can be empty when no codes were identified, including when the input has no audio stream. These codes do not indicate available translations.
	InputLanguageCodes []string `json:"inputLanguageCodes,omitempty"`
	// Credits extracted from the video's on-screen text.
	Credits *Credits `json:"credits,omitempty"`
	// Information about the analysis that generated this result. (required)
	Metadata *Metadata `json:"metadata,omitempty"`
}
