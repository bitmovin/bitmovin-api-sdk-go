package model

// SceneAnalysisDetailsResponse model
type SceneAnalysisDetailsResponse struct {
	Scenes      []Scene  `json:"scenes,omitempty"`
	Duration    *float64 `json:"duration,omitempty"`
	Description *string  `json:"description,omitempty"`
	// Inferred title representing the analyzed content as a whole. If omitted or null, the title is not available.
	Title                       *string   `json:"title,omitempty"`
	Keywords                    []string  `json:"keywords,omitempty"`
	Ratings                     []Rating  `json:"ratings,omitempty"`
	SensitiveTopics             []string  `json:"sensitiveTopics,omitempty"`
	IabSensitiveTopicTaxonomies []string  `json:"iabSensitiveTopicTaxonomies,omitempty"`
	InputLanguageCodes          []string  `json:"inputLanguageCodes,omitempty"`
	Credits                     *Credits  `json:"credits,omitempty"`
	Metadata                    *Metadata `json:"metadata,omitempty"`
}
