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
	// A detailed breakdown of individual camera shots within this scene, providing granular analysis of visual content and subjects
	Shots []Shot `json:"shots,omitempty"`
}
