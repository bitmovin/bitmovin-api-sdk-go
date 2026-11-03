package model

// Scene model
type Scene struct {
	// Short AI-generated title describing the scene.
	Title *string `json:"title,omitempty"`
	// Start time of the scene in seconds from the beginning of the video.
	StartInSeconds *float64 `json:"startInSeconds,omitempty"`
	// End time of the scene in seconds from the beginning of the video.
	EndInSeconds *float64 `json:"endInSeconds,omitempty"`
	// Identifier assigned to this scene within the analysis result.
	Id *string `json:"id,omitempty"`
	// Structured description of this scene.
	Content *Content `json:"content,omitempty"`
	// Concise summary of the scene.
	Summary *string `json:"summary,omitempty"`
	// Detailed description of the scene, providing more context than summary.
	VerboseSummary *string `json:"verboseSummary,omitempty"`
	// Sensitive topics identified in this scene, expressed as free text rather than taxonomy identifiers. (required)
	SensitiveTopics []string `json:"sensitiveTopics,omitempty"`
	// Free-text keywords describing this scene. Preserve the returned spelling and case. (required)
	Keywords []string `json:"keywords,omitempty"`
	// IAB classification of this scene.
	Iab *IabTaxonomy `json:"iab,omitempty"`
	// The detected type of scene based on content analysis
	Type SceneType `json:"type,omitempty"`
	// Camera shots with available analysis associated with this scene. The list can be empty or omitted and need not include every shot.
	Shots []Shot `json:"shots,omitempty"`
	// Suitability assessment for an ad break at the end of this scene. This does not confirm that an ad was placed.
	AdOpportunityInformation *AdOpportunity `json:"adOpportunityInformation,omitempty"`
}
