package model

// Scene model
type Scene struct {
	StartInSeconds  *float32     `json:"startInSeconds,omitempty"`
	EndInSeconds    *float32     `json:"endInSeconds,omitempty"`
	Id              *string      `json:"id,omitempty"`
	Content         *Content     `json:"content,omitempty"`
	Summary         *string      `json:"summary,omitempty"`
	SensitiveTopics []string     `json:"sensitiveTopics,omitempty"`
	Keywords        []string     `json:"keywords,omitempty"`
	Iab             *IabTaxonomy `json:"iab,omitempty"`
}
