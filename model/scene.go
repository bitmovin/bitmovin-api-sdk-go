package model

// Scene model
type Scene struct {
	Title           *string      `json:"title,omitempty"`
	StartInSeconds  *float32     `json:"startInSeconds,omitempty"`
	EndInSeconds    *float32     `json:"endInSeconds,omitempty"`
	Id              *string      `json:"id,omitempty"`
	Content         *Content     `json:"content,omitempty"`
	Summary         *string      `json:"summary,omitempty"`
	VerboseSummary  *string      `json:"verboseSummary,omitempty"`
	SensitiveTopics []string     `json:"sensitiveTopics,omitempty"`
	Keywords        []string     `json:"keywords,omitempty"`
	Iab             *IabTaxonomy `json:"iab,omitempty"`
	// The detected type of scene based on content analysis
	Type SceneType `json:"type,omitempty"`
	// Confidence score for the detected scene type (0.0 to 1.0)
	TypeConfidence *float64 `json:"typeConfidence,omitempty"`
}
